// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build linux

package tshttpproxy

import (
	"net/http"
	"net/url"

	"github.com/WebP2P/dexnet/feature/buildfeatures"
	"github.com/WebP2P/dexnet/version/distro"
)

func init() {
	sysProxyFromEnv = linuxSysProxyFromEnv
}

func linuxSysProxyFromEnv(req *http.Request) (*url.URL, error) {
	if buildfeatures.HasSynology && distro.Get() == distro.Synology {
		return synologyProxyFromConfigCached(req)
	}
	return nil, nil
}
