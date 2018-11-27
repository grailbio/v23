// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build linux darwin windows

// Package v23 implements a RuntimeFactory suitable for a variety of network
// configurations appropriate for grail. Roaming in particular is disabled
// to avoid bringing in any cgo dependencies.
//
package v23

import (
	"v.io/v23/flow"
	"v.io/x/ref/runtime/factories/library"
	"v.io/x/ref/runtime/protocols/lib/websocket"
	_ "v.io/x/ref/runtime/protocols/tcp" // Initialize tcp, ws and wsh.
	_ "v.io/x/ref/runtime/protocols/ws"
	_ "v.io/x/ref/runtime/protocols/wsh"
)

func init() {
	library.Roam = false
	library.CloudVM = true
	library.ConfigureLoggingFromFlags = true
	library.ReservedNameDispatcher = true
	flow.RegisterUnknownProtocol("wsh", websocket.WSH{})
	library.EnableCommandlineFlags()
}
