// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build cgo || !darwin

// systray is a minimal Tailscale systray application.
package main

import (
	"flag"

	"github.com/WebP2P/dexnet/client/local"
	"github.com/WebP2P/dexnet/client/systray"
	"github.com/WebP2P/dexnet/paths"
)

var socket = flag.String("socket", paths.DefaultTailscaledSocket(), "path to tailscaled socket")

func main() {
	flag.Parse()
	lc := &local.Client{Socket: *socket}
	new(systray.Menu).Run(lc)
}
