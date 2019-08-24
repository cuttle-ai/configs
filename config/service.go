// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package config

/* This file contains the service interface defnitions where a service is the one that require configuration */

//Service is the interface to be implemented for representing a service that requires configuration
type Service interface {
	//Init will init the configuration for the service
	Init(Config) error
}
