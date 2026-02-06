// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !ts_omit_webclient

package main

import (
	"github.com/WebP2P/dexnet/client/local"
	"github.com/WebP2P/dexnet/ipn/ipnlocal"
	"github.com/WebP2P/dexnet/paths"
)

func init() {
	hookConfigureWebClient.Set(func(lb *ipnlocal.LocalBackend) {
		lb.ConfigureWebClient(&local.Client{
			Socket:        args.socketpath,
			UseSocketOnly: args.socketpath != paths.DefaultTailscaledSocket(),
		})
	})
}
