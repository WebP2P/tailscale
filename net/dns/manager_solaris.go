// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package dns

import (
	"github.com/WebP2P/dexnet/control/controlknobs"
	"github.com/WebP2P/dexnet/health"
	"github.com/WebP2P/dexnet/types/logger"
	"github.com/WebP2P/dexnet/util/eventbus"
	"github.com/WebP2P/dexnet/util/syspolicy/policyclient"
)

func NewOSConfigurator(logf logger.Logf, health *health.Tracker, bus *eventbus.Bus, _ policyclient.Client, _ *controlknobs.Knobs, iface string) (OSConfigurator, error) {
	return newDirectManager(logf, health, bus), nil
}
