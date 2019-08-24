// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package version has the version information about the application
package version

import "fmt"

/* This file contains the app version definitions */

//Version denotes the version
type Version struct {
	//Major is the major version version
	Major int
	//Minor is the minor version number
	Minor int
	//Patches is the bug fixes number
	Patches int
}

//String is the stringer implementation for Version
func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patches)
}

var (
	//V1 is the version 1 of the application
	V1 = Version{1, 0, 0}
)

var (
	//Default is the default version of the application
	Default = V1
)
