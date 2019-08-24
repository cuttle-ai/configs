// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

/* This file contains the config store interface defnitions */

//Store stores the configuration required for an application in a remote/local server
type Store interface {
	//GetConfig returns the configuration from the store
	//config should be the name of the config
	GetConfig(config string) (Config, error)
}
