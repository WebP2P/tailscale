// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"testing"

	"github.com/WebP2P/dexnet/tstest/deptest"
)

func TestDeps(t *testing.T) {
	deptest.DepChecker{
		BadDeps: map[string]string{
			"testing":                            "do not use testing package in production code",
			"gvisor.dev/gvisor/pkg/buffer":       "https://github.com/tailscale/tailscale/issues/9756",
			"gvisor.dev/gvisor/pkg/cpuid":        "https://github.com/tailscale/tailscale/issues/9756",
			"gvisor.dev/gvisor/pkg/tcpip":        "https://github.com/tailscale/tailscale/issues/9756",
			"gvisor.dev/gvisor/pkg/tcpip/header": "https://github.com/tailscale/tailscale/issues/9756",
			"github.com/WebP2P/dexnet/wgengine/filter":      "brings in bart, etc",
			"github.com/bits-and-blooms/bitset":  "unneeded in CLI",
			"github.com/WebP2P/dexnet/net/ipset":            "unneeded in CLI",
		},
	}.Check(t)
}
