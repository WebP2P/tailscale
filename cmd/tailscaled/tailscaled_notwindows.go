// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

//go:build !windows && go1.19

package main // import "github.com/WebP2P/dexnet/cmd/tailscaled"

import "github.com/WebP2P/dexnet/logpolicy"

func isWindowsService() bool { return false }

func runWindowsService(pol *logpolicy.Policy) error { panic("unreachable") }

func beWindowsSubprocess() bool { return false }
