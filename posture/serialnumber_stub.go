// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// js: not implemented
// plan9: not implemented
// solaris: currently unsupported by go-smbios:
// https://github.com/digitalocean/go-smbios/pull/21

//go:build solaris || plan9 || js || wasm || tamago || aix || (darwin && !cgo && !ios)

package posture

import (
	"errors"

	"github.com/WebP2P/dexnet/types/logger"
	"github.com/WebP2P/dexnet/util/syspolicy/policyclient"
)

// GetSerialNumber returns client machine serial number(s).
func GetSerialNumbers(polc policyclient.Client, _ logger.Logf) ([]string, error) {
	return nil, errors.New("not implemented")
}
