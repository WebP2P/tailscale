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
			"github.com/WebP2P/dexnet/tailcfg": "circular dependency via go generate",
			"github.com/WebP2P/dexnet/version": "circular dependency via go generate",
		},
	}.Check(t)
}
