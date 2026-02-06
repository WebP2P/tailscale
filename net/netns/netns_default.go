// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !linux && !windows && !darwin

package netns

import (
	"syscall"

	"github.com/WebP2P/dexnet/net/netmon"
	"github.com/WebP2P/dexnet/types/logger"
)

func control(logger.Logf, *netmon.Monitor) func(network, address string, c syscall.RawConn) error {
	return controlC
}

// controlC does nothing to c.
func controlC(network, address string, c syscall.RawConn) error {
	return nil
}

func UseSocketMark() bool {
	return false
}
