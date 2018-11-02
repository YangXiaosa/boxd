// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"math/rand"
	"time"

	"github.com/BOXFoundation/boxd/log"
)

const (
	workDir   = "../.devconfig/"
	keyDir    = "../keyfile/"
	walletDir = "../.devconfig/ws1/box_keystore"

	testPassphrase = "zaq12wsx"

	rpcAddr = "127.0.0.1:19191"

	dockerComposeFile = "../docker-compose.yml"
)

var logger = log.NewLogger("tests") // logger

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	txTest()
}

func txTest() {
	nodeCount := 3
	testCount := 2

	// prepare environment and clean history data
	//if err := prepareEnv(nodeCount); err != nil {
	//	logger.Fatal(err)
	//}

	// start nodes
	if err := startNodes(); err != nil {
		logger.Fatal(err)
	}

	// wait for nodes to be ready
	logger.Info("wait node running for 12 seconds")
	time.Sleep(12 * time.Second)

	// get addresses of three miners
	var minersAddr []string
	for i := 0; i < nodeCount; i++ {
		addr, err := minerAddress(i)
		if err != nil {
			logger.Fatal(err)
		}
		minersAddr = append(minersAddr, addr)
	}
	logger.Infof("minersAddr: %v", minersAddr)

	//  generate addresses of test accounts(T1, T2)
	var testsAddr []string
	for i := 0; i < testCount; i++ {
		addr, err := newAccount()
		if err != nil {
			logger.Fatal(err)
		}
		testsAddr = append(testsAddr, addr)
	}
	logger.Infof("testsAddr: %v", testsAddr)
	defer removeNewKeystoreFiles()

	// wait for some blocks to generate
	//logger.Info("wait mining for 5 seconds")
	//time.Sleep(5 * time.Second)

	// get balance of miners
	logger.Info("start getting balance of miners")
	var minersBalance []uint64
	for i := 0; i < nodeCount; i++ {
		amount, err := balanceFor(minersAddr[i], rpcAddr)
		if err != nil {
			logger.Fatal(err)
		}
		minersBalance = append(minersBalance, amount)
	}
	logger.Infof("minersBalance: %v", minersBalance)

	//// wait some time
	//time.Sleep(15 * time.Second)

	// check balance of each miners

	// check miners' balances

	// get initial balance of test accounts
	logger.Info("start getting balance of test accounts")
	var testsBalance []uint64
	for i := 0; i < testCount; i++ {
		amount, err := balanceFor(testsAddr[i], rpcAddr)
		if err != nil {
			logger.Fatal(err)
		}
		testsBalance = append(testsBalance, amount)
	}
	logger.Infof("testsBalance: %v", testsBalance)

	// create a transaction and execute it
	//accT1 := testsAddr[0]
	txAmount := uint64(10000 + rand.Intn(10000))
	if err := execTx(minersAddr[0], minersAddr[1], txAmount, rpcAddr); err != nil {
		logger.Fatal(err)
	}
	logger.Infof("have sent %d from %s to %s", txAmount, minersAddr[0], minersAddr[1])

	// wait transaction brought to chain
	broadcastTime := 10 * time.Second
	logger.Infof("wait transaction broadcasted for %v", broadcastTime)
	time.Sleep(broadcastTime)

	// check the balance of T1 on all nodes
	amount, err := balanceFor(minersAddr[0], rpcAddr)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("balance for account %s is %v", minersAddr[0], amount)
	amount, err = balanceFor(minersAddr[1], rpcAddr)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infof("balance for account %s is %v", minersAddr[1], amount)

	//// transfer 2424 from T1 to T2
	//someBoxes = 1000 + rand.Intn(1000)

	//// wait some time
	//time.Sleep(15 * time.Second)

	// check the balance of M1 and T1 on all nodes

	// test finished, stopNodes
	if err := stopNodes(); err != nil {
		logger.Fatal(err)
	}

	// clear databases and logs
	//if err := tearDown(nodeCount); err != nil {
	//	logger.Fatal(err)
	//}
}
