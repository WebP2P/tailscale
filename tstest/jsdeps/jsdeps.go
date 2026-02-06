// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package jsdeps is a just a list of the packages we import in the
// JavaScript/WASM build, to let us test that our transitive closure of
// dependencies doesn't accidentally grow too large, since binary size
// is more of a concern.
package jsdeps

import (
	_ "bytes"
	_ "context"
	_ "encoding/hex"
	_ "encoding/json"
	_ "fmt"
	_ "log"
	_ "math/rand/v2"
	_ "net"
	_ "strings"
	_ "time"

	_ "golang.org/x/crypto/ssh"
	_ "github.com/WebP2P/dexnet/control/controlclient"
	_ "github.com/WebP2P/dexnet/ipn"
	_ "github.com/WebP2P/dexnet/ipn/ipnserver"
	_ "github.com/WebP2P/dexnet/net/netaddr"
	_ "github.com/WebP2P/dexnet/net/netns"
	_ "github.com/WebP2P/dexnet/net/tsdial"
	_ "github.com/WebP2P/dexnet/safesocket"
	_ "github.com/WebP2P/dexnet/tailcfg"
	_ "github.com/WebP2P/dexnet/types/logger"
	_ "github.com/WebP2P/dexnet/wgengine"
	_ "github.com/WebP2P/dexnet/wgengine/netstack"
	_ "github.com/WebP2P/dexnet/words"
)
