// Copyright 2019 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package system

import (
	"runtime"
	"testing"
)

func TestExecWinshell(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.SkipNow()
	}
	for _, v := range listWinShell {
		_, err := ExecWinshell(v, false, `dir C:\`)
		if err != nil {
			t.Fatal(err)
		}
	}
}