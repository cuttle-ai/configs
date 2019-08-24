// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/hashicorp/vault/api"
)

/* This file has the definitions of vault store */

const (
	//VAULT_ADDRESS is the environement variable for address of vault server
	VAULT_ADDRESS = "CUTTLE_AI_CONFIG_VAULT_ADDRESS"
	//DEFAULT_VAULT_ADDRESS is the environment variable for address of vault server
	DEFAULT_VAULT_ADDRESS = "VAULT_ADDR"
	//VAULT_TOKEN is the environment variable for vault access token
	VAULT_TOKEN = "CUTTLE_AI_CONFIG_VAULT_TOKEN"
	//VAULT_PATH is the environement variable for vault path
	VAULT_PATH = "CUTTLE_AI_CONFIG_VAULT_DEFAULT_PATH"
)

//Vault is config store that stores the configuration in hashcopr vault
//It fetches the configuration from vault on demand
type Vault struct {
	//Config has the configuration required for connecting with the Vault server
	Config *api.Config
	//Token to commuincate with the vault server
	Token string
	//Path is the path from which the secrets has to be fetched
	Path string
}

//NewVault will create a new vault and inits the configuration
//If any error found while initing the config, it will return it with nil vault
//Else nil error is returned.
func NewVault() (*Vault, error) {
	v := &Vault{}
	err := v.InitConfig()
	if err != nil {
		return nil, err
	}
	return v, nil
}

//GetConfig returns the config along with the error if any while getting the configuration from
//remote/local vault server. If no error happens it will return non-nil Config and nil as error.
//If error happens it will return the error and nil as config
func (v *Vault) GetConfig(config string) (Config, error) {
	/*
	 * We will init config if the vault configuration is nil
	 * We will first create a client and set the token to communicate with the vault
	 * Then we will read the configuration metadata
	 * Then we read the configuration data
	 * Will write it in config
	 */
	//initing the vault configuration if not existing
	if v.Config == nil || len(v.Token) == 0 {
		err := v.InitConfig()
		if err != nil {
			return nil, err
		}
	}

	//getting the client to communicate with the vault
	client, err := api.NewClient(v.Config)
	if err != nil {
		fmt.Println("Error while getting a new vault client")
		return nil, err
	}
	client.SetToken(v.Token)

	//reading the configuration metadata
	c, err := client.Logical().ReadWithData(v.Path+"/metadata/"+config, nil)
	if err != nil {
		fmt.Println("Error while reading the secrets from", v.Path)
		return nil, err
	}
	ver, err := c.Data["current_version"].(json.Number).Int64()
	if err != nil {
		//Error while getting the version of secrets
		fmt.Println("Error while getting the version of the config", config)
		return nil, err
	}

	//reading the configuration data
	secrets, err := client.Logical().ReadWithData(v.Path+"/data/"+config, map[string][]string{
		"version": []string{strconv.Itoa(int(ver))},
	})
	if err != nil {
		fmt.Println("Error while reading the secrets from", v.Path)
		return nil, err
	}

	//writing the secrets in config
	conf := secrets.Data["data"].(map[string]interface{})
	configObj := Config{}
	for k, v := range conf {
		configObj[k] = v.(string)
	}

	return configObj, nil
}

//InitConfig will init the configuration from environment variables
//The only supported method of providing configuration is environment variables.
func (v *Vault) InitConfig() error {
	/*
	 * We will init the config
	 * We will get the address from the environment variable
	 * We will get the token from the environment variable
	 * We will get the secret path
	 * Finally we will set the token, path and config
	 */

	//initing the config
	config := new(api.Config)

	//getting the vault address. We will try both sensibull and default vault address
	config.Address = os.Getenv(VAULT_ADDRESS)
	if len(config.Address) == 0 {
		config.Address = os.Getenv(DEFAULT_VAULT_ADDRESS)
	}
	if len(config.Address) == 0 {
		//checking whether the vault address is empty
		return errors.New("couldn't find the vault address. Please set that in either of the following environment variables :- " + VAULT_ADDRESS + ", " + DEFAULT_VAULT_ADDRESS)
	}

	//getting the token
	token := os.Getenv(VAULT_TOKEN)
	if len(token) == 0 {
		//checking whether the vault token is empty
		return errors.New("Couldn't find the vault access token. Please set the token in " + VAULT_TOKEN)
	}

	//getting the path
	secretPath := os.Getenv(VAULT_PATH)
	if len(secretPath) == 0 {
		//checking whether the vault secrets path is empty
		return errors.New("Couldn't find the vault secrets path. Please set the path in " + VAULT_PATH)
	}

	//setting the config and token
	v.Config = config
	v.Token = token
	v.Path = secretPath

	return nil
}
