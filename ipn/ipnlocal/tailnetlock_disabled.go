// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build ts_omit_tailnetlock

package ipnlocal

import (
	"github.com/WebP2P/dexnet/ipn"
	"github.com/WebP2P/dexnet/ipn/ipnstate"
	"github.com/WebP2P/dexnet/tka"
	"github.com/WebP2P/dexnet/types/netmap"
)

type tkaState struct {
	authority *tka.Authority
}

func (b *LocalBackend) initTKALocked() error {
	return nil
}

func (b *LocalBackend) tkaSyncIfNeeded(nm *netmap.NetworkMap, prefs ipn.PrefsView) error {
	return nil
}

func (b *LocalBackend) tkaFilterNetmapLocked(nm *netmap.NetworkMap) {}

func (b *LocalBackend) NetworkLockStatus() *ipnstate.NetworkLockStatus {
	return &ipnstate.NetworkLockStatus{Enabled: false}
}
