// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package portmapper registers support for NAT-PMP, PCP, and UPnP port
// mapping protocols to help get direction connections through NATs.
package portmapper

import (
	"github.com/WebP2P/dexnet/feature"
	"github.com/WebP2P/dexnet/net/netmon"
	"github.com/WebP2P/dexnet/net/portmapper"
	"github.com/WebP2P/dexnet/net/portmapper/portmappertype"
	"github.com/WebP2P/dexnet/types/logger"
	"github.com/WebP2P/dexnet/util/eventbus"
)

func init() {
	feature.Register("portmapper")
	portmappertype.HookNewPortMapper.Set(newPortMapper)
}

func newPortMapper(
	logf logger.Logf,
	bus *eventbus.Bus,
	netMon *netmon.Monitor,
	disableUPnPOrNil func() bool,
	onlyTCP443OrNil func() bool) portmappertype.Client {

	pm := portmapper.NewClient(portmapper.Config{
		EventBus: bus,
		Logf:     logf,
		NetMon:   netMon,
		DebugKnobs: &portmapper.DebugKnobs{
			DisableAll:      onlyTCP443OrNil,
			DisableUPnPFunc: disableUPnPOrNil,
		},
	})
	pm.SetGatewayLookupFunc(netMon.GatewayAndSelfIP)
	return pm
}
