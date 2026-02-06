// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package iosdeps is a just a list of the packages we import on iOS, to let us
// test that our transitive closure of dependencies on iOS doesn't accidentally
// grow too large, as we've historically been memory constrained there.
package iosdeps

import (
	_ "bufio"
	_ "bytes"
	_ "context"
	_ "crypto/rand"
	_ "crypto/sha256"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "io/fs"
	_ "log"
	_ "math"
	_ "net"
	_ "net/http"
	_ "os"
	_ "os/signal"
	_ "path/filepath"
	_ "runtime"
	_ "runtime/debug"
	_ "strings"
	_ "sync"
	_ "sync/atomic"
	_ "syscall"
	_ "time"
	_ "unsafe"

	_ "github.com/tailscale/wireguard-go/device"
	_ "github.com/tailscale/wireguard-go/tun"
	_ "go4.org/mem"
	_ "golang.org/x/sys/unix"
	_ "github.com/WebP2P/dexnet/hostinfo"
	_ "github.com/WebP2P/dexnet/ipn"
	_ "github.com/WebP2P/dexnet/ipn/ipnlocal"
	_ "github.com/WebP2P/dexnet/ipn/localapi"
	_ "github.com/WebP2P/dexnet/logtail"
	_ "github.com/WebP2P/dexnet/logtail/filch"
	_ "github.com/WebP2P/dexnet/net/dns"
	_ "github.com/WebP2P/dexnet/net/netaddr"
	_ "github.com/WebP2P/dexnet/net/tsdial"
	_ "github.com/WebP2P/dexnet/net/tstun"
	_ "github.com/WebP2P/dexnet/paths"
	_ "github.com/WebP2P/dexnet/types/empty"
	_ "github.com/WebP2P/dexnet/types/logger"
	_ "github.com/WebP2P/dexnet/util/clientmetric"
	_ "github.com/WebP2P/dexnet/util/dnsname"
	_ "github.com/WebP2P/dexnet/version"
	_ "github.com/WebP2P/dexnet/wgengine"
	_ "github.com/WebP2P/dexnet/wgengine/router"
)
