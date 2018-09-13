// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package node

import (
	"net"
	"os"
	"runtime"

	"github.com/BOXFoundation/Quicksilver/log"
	p2p "github.com/BOXFoundation/Quicksilver/p2p"
	"github.com/jbenet/goprocess"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/spf13/viper"
)

// RootProcess is the root process of the app
var RootProcess goprocess.Process

var logger *log.Logger

func init() {
	RootProcess = goprocess.WithSignals(os.Interrupt)
	logger = log.NewLogger("node")
}

// Start function starts node server.
func Start(v *viper.Viper) error {
	runtime.GOMAXPROCS(runtime.NumCPU())

	log.Setup(v) // setup logger

	var host, err = p2p.NewDefaultHost(RootProcess, net.ParseIP(v.GetString("node.listen.address")), uint(v.GetInt("node.listen.port")))
	if err != nil {
		logger.Error(err)
		return err
	}

	// connect to other peers passed via commandline
	for _, addr := range v.GetStringSlice("node.addpeer") {
		if maddr, err := ma.NewMultiaddr(addr); err == nil {
			err := host.ConnectPeer(RootProcess, maddr)
			if err != nil {
				logger.Warn(err)
			} else {
				logger.Infof("Peer %s connected.", maddr)
			}
		} else {
			logger.Warnf("Invalid multiaddress %s", addr)
		}
	}

	select {
	case <-RootProcess.Closing():
		logger.Info("Box server is shutting down...")
	}

	select {
	case <-RootProcess.Closed():
		logger.Info("Box server is down.")
	}

	return nil
}