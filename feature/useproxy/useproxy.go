// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package useproxy registers support for using system proxies.
package useproxy

import (
	"github.com/WebP2P/dexnet/feature"
	"github.com/WebP2P/dexnet/net/tshttpproxy"
)

func init() {
	feature.HookProxyFromEnvironment.Set(tshttpproxy.ProxyFromEnvironment)
	feature.HookProxyInvalidateCache.Set(tshttpproxy.InvalidateCache)
	feature.HookProxyGetAuthHeader.Set(tshttpproxy.GetAuthHeader)
	feature.HookProxySetSelfProxy.Set(tshttpproxy.SetSelfProxy)
	feature.HookProxySetTransportGetProxyConnectHeader.Set(tshttpproxy.SetTransportGetProxyConnectHeader)
}
