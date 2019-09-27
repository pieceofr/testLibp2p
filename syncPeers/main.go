package main

import (
	"flag"
	"fmt"
	"time"
)

var servant SimpleNode
var client SimpleNode

var runOpt int
var runPort int
var runKey string
var runAddr string

func main() {

	flag.IntVar(&runOpt, "type", 0, "0:servant 1:client")
	flag.IntVar(&runPort, "port", 12140, "port to run node")
	flag.StringVar(&runKey, "key", "./peer.private", "node private key path")
	flag.StringVar(&runAddr, "addr", "", "Connect to Address")

	flag.Parse()
	switch runOpt {
	case 0:
		runServant(runPort)
	case 1:
	default:
	}

	for {
		time.Sleep(10 * time.Second)
	}
}

func runServant(port int) {
	servant.createServHost(0, port, "127.0.0.1", runKey)
	servant.Listen()
	if runAddr != "" {
		servant.ConnectTo(runAddr)
	}
}

func runClient(addr string) {
	//client1.CreateNode(1, 0, "")
	client.createServHost(1, 0, "127.0.0.1", runKey)
	if runAddr != "" {
		client.ConnectTo(addr)
	} else {
		fmt.Println("Client does not know where to connect")
	}
}
