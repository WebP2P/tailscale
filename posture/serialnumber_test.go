// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package posture

import (
	"testing"

	"github.com/WebP2P/dexnet/types/logger"
	"github.com/WebP2P/dexnet/util/syspolicy/policyclient"
)

func TestGetSerialNumber(t *testing.T) {
	// ensure GetSerialNumbers is implemented
	// or covered by a stub on a given platform.
	_, _ = GetSerialNumbers(policyclient.NoPolicyClient{}, logger.Discard)
}
