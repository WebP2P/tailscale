// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package ace registers support for Alternate Connectivity Endpoints (ACE).
package ace

import (
	"net/netip"

	"github.com/WebP2P/dexnet/control/controlhttp"
	"github.com/WebP2P/dexnet/net/ace"
	"github.com/WebP2P/dexnet/net/netx"
)

func init() {
	controlhttp.HookMakeACEDialer.Set(mkDialer)
}

func mkDialer(dialer netx.DialFunc, aceHost string, optIP netip.Addr) netx.DialFunc {
	return (&ace.Dialer{
		ACEHost:   aceHost,
		ACEHostIP: optIP, // may be zero
		NetDialer: dialer,
	}).Dial
}
