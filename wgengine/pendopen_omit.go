// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ts_omit_debug

package wgengine

import (
	"github.com/WebP2P/dexnet/net/packet"
	"github.com/WebP2P/dexnet/net/tstun"
	"github.com/WebP2P/dexnet/wgengine/filter"
)

type flowtrackTuple = struct{}

type pendingOpenFlow struct{}

func (*userspaceEngine) trackOpenPreFilterIn(pp *packet.Parsed, t *tstun.Wrapper) (res filter.Response) {
	panic("unreachable")
}

func (*userspaceEngine) trackOpenPostFilterOut(pp *packet.Parsed, t *tstun.Wrapper) (res filter.Response) {
	panic("unreachable")
}
