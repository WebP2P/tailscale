// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows

package controlclient

import (
	"github.com/WebP2P/dexnet/tailcfg"
	"github.com/WebP2P/dexnet/types/key"
	"github.com/WebP2P/dexnet/util/syspolicy/policyclient"
)

// signRegisterRequest on non-supported platforms always returns errNoCertStore.
func signRegisterRequest(polc policyclient.Client, req *tailcfg.RegisterRequest, serverURL string, serverPubKey, machinePubKey key.MachinePublic) error {
	return errNoCertStore
}
