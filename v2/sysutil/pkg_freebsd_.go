// Copyright 2021 Jonas mg
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// System: FreeBSD

package sysutil

import (
	"github.com/tredoe/osutil/v2/executil"
	"github.com/tredoe/osutil/v2/internal"
)

const (
	filePkg = "pkg"
	pathPkg = "/usr/sbin/pkg"
)

// ManagerPkg is the interface to handle the FreeBSD package manager,
// called 'package' or 'pkg'.
type ManagerPkg struct {
	pathExec string
	sudo     string
	cmd      *executil.Command
}

// NewManagerPkg returns the Pkg package manager.
func NewManagerPkg() ManagerPkg {
	return ManagerPkg{
		pathExec: pathPkg,
		sudo:     "/usr/local/bin/sudo",
		cmd:      excmd.Command("", "").BadExitCodes([]int{1}),
	}
}

func (m ManagerPkg) setExecPath(p string) { m.pathExec = p }

func (m ManagerPkg) Cmd() *executil.Command { return m.cmd }

func (m ManagerPkg) ExecPath() string { return m.pathExec }

func (m ManagerPkg) PackageType() string { return Pkg.String() }

func (m ManagerPkg) Install(name ...string) error {
	internal.Log.Print(taskInstall)
	args := append([]string{pathPkg, "install", "-y"}, name...)

	_, err := m.cmd.Command(m.sudo, args...).Run()
	return err
}

func (m ManagerPkg) Remove(name ...string) error {
	internal.Log.Print(taskRemove)
	args := append([]string{pathPkg, "delete", "-y"}, name...)

	_, err := m.cmd.Command(m.sudo, args...).Run()
	return err
}

func (m ManagerPkg) Purge(name ...string) error {
	internal.Log.Print(taskPurge)
	return m.Remove(name...)
}

func (m ManagerPkg) Update() error {
	internal.Log.Print(taskUpdate)
	_, err := m.cmd.Command(m.sudo, pathPkg, "update").Run()
	return err
}

func (m ManagerPkg) Upgrade() error {
	internal.Log.Print(taskUpgrade)
	_, err := m.cmd.Command(m.sudo, pathPkg, "upgrade", "-y").Run()
	return err
}

func (m ManagerPkg) Clean() error {
	internal.Log.Print(taskClean)
	_, err := m.cmd.Command(m.sudo, pathPkg, "autoremove", "-y").Run()
	if err != nil {
		return err
	}
	_, err = m.cmd.Command(m.sudo, pathPkg, "clean", "-y").Run()
	return err
}

func (m ManagerPkg) ImportKey(alias, keyUrl string) error {
	return ErrManagCmd
}

func (m ManagerPkg) ImportKeyFromServer(alias, keyServer, key string) error {
	return ErrManagCmd
}

func (m ManagerPkg) RemoveKey(alias string) error {
	return ErrManagCmd
}

func (m ManagerPkg) AddRepo(alias string, url ...string) error {
	return ErrManagCmd
}

func (m ManagerPkg) RemoveRepo(r string) error {
	return ErrManagCmd
}
