// Copyright 2021 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package osutil

import (
	"os"

	"github.com/tredoe/osutil/executil"
)

var excmd = executil.NewCommand("", "").
	Stdout(os.Stdout).Stderr(os.Stderr).
	Env(append([]string{"LANG=C"}, os.Environ()...))
