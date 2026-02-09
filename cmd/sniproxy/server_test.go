// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

package main

import (
	"net/netip"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/WebP2P/dexnet/tailcfg"
	"github.com/WebP2P/dexnet/types/appctype"
)

func TestMakeConnectorsFromConfig(t *testing.T) {
	tcs := []struct {
		name  string
		input *appctype.AppConnectorConfig
		want  map[appctype.ConfigID]connector
	}{
		{
			"empty",
			&appctype.AppConnectorConfig{},
			nil,
		},
		{
			"DNAT",
			&appctype.AppConnectorConfig{
				DNAT: map[appctype.ConfigID]appctype.DNATConfig{
					"swiggity_swooty": {
						Addrs: []netip.Addr{netip.MustParseAddr("100.64.0.1"), netip.MustParseAddr("fd0d:e100:d3c5::1")},
						To:    []string{"example.org"},
						IP:    []tailcfg.ProtoPortRange{{Proto: 0, Ports: tailcfg.PortRange{First: 0, Last: 65535}}},
					},
				},
			},
			map[appctype.ConfigID]connector{
				"swiggity_swooty": {
					Handlers: map[target]handler{
						{
							Dest:     netip.MustParsePrefix("100.64.0.1/32"),
							Matching: tailcfg.ProtoPortRange{Proto: 0, Ports: tailcfg.PortRange{First: 0, Last: 65535}},
						}: &tcpRoundRobinHandler{To: []string{"example.org"}, ReachableIPs: []netip.Addr{netip.MustParseAddr("100.64.0.1"), netip.MustParseAddr("fd0d:e100:d3c5::1")}},
						{
							Dest:     netip.MustParsePrefix("fd0d:e100:d3c5::1/128"),
							Matching: tailcfg.ProtoPortRange{Proto: 0, Ports: tailcfg.PortRange{First: 0, Last: 65535}},
						}: &tcpRoundRobinHandler{To: []string{"example.org"}, ReachableIPs: []netip.Addr{netip.MustParseAddr("100.64.0.1"), netip.MustParseAddr("fd0d:e100:d3c5::1")}},
					},
				},
			},
		},
		{
			"SNIProxy",
			&appctype.AppConnectorConfig{
				SNIProxy: map[appctype.ConfigID]appctype.SNIProxyConfig{
					"swiggity_swooty": {
						Addrs:          []netip.Addr{netip.MustParseAddr("100.64.0.1"), netip.MustParseAddr("fd0d:e100:d3c5::1")},
						AllowedDomains: []string{"example.org"},
						IP:             []tailcfg.ProtoPortRange{{Proto: 0, Ports: tailcfg.PortRange{First: 0, Last: 65535}}},
					},
				},
			},
			map[appctype.ConfigID]connector{
				"swiggity_swooty": {
					Handlers: map[target]handler{
						{
							Dest:     netip.MustParsePrefix("100.64.0.1/32"),
							Matching: tailcfg.ProtoPortRange{Proto: 0, Ports: tailcfg.PortRange{First: 0, Last: 65535}},
						}: &tcpSNIHandler{Allowlist: []string{"example.org"}, ReachableIPs: []netip.Addr{netip.MustParseAddr("100.64.0.1"), netip.MustParseAddr("fd0d:e100:d3c5::1")}},
						{
							Dest:     netip.MustParsePrefix("fd0d:e100:d3c5::1/128"),
							Matching: tailcfg.ProtoPortRange{Proto: 0, Ports: tailcfg.PortRange{First: 0, Last: 65535}},
						}: &tcpSNIHandler{Allowlist: []string{"example.org"}, ReachableIPs: []netip.Addr{netip.MustParseAddr("100.64.0.1"), netip.MustParseAddr("fd0d:e100:d3c5::1")}},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			connectors := makeConnectorsFromConfig(tc.input)

			if diff := cmp.Diff(connectors, tc.want,
				cmpopts.IgnoreFields(tcpRoundRobinHandler{}, "DialContext"),
				cmpopts.IgnoreFields(tcpSNIHandler{}, "DialContext"),
				cmp.Comparer(func(x, y netip.Addr) bool {
					return x == y
				})); diff != "" {
				t.Fatalf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
