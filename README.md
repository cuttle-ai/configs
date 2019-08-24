# Configs

Can fetch the configuration from secret management services like vault.
Currently we support vault

## Usage

```go
func main() {
	v, err := config.NewVault()
	checkError(err)
	config, err := v.GetConfig("auth-service")
	checkError(err)
	fmt.Println(config)
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

It do require the following environment variables to work

```
export CUTTLE_AI_CONFIG_VAULT_ADDRESS='vault.cuttle.ai'
export CUTTLE_AI_CONFIG_VAULT_TOKEN='<token-provied-to-access-vault>'
export CUTTLE_AI_CONFIG_VAULT_DEFAULT_PATH='cuttle-ai-development'
```
