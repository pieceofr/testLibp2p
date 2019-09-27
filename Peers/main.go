package main

import (
	"flag"
	"fmt"
	"time"
)

var servant1 SimpleNode
var servant2 SimpleNode
var client1 SimpleNode
var client2 SimpleNode

func main() {

	var runOpt int
	var addr string
	flag.IntVar(&runOpt, "run", 1, "run which setting")
	flag.StringVar(&addr, "addr", "", "Address to Connect")
	flag.Parse()
	switch runOpt {
	case 1:
		go runServant1()
		fmt.Println("===Start Run Servant1===")
	case 2:
		go runServant2()
		fmt.Println("===Start Run Servant2===")
	case 3:
		go runClient1(addr)
		fmt.Println("===Start Run runClient1===")
	case 4:
		go runClient2(addr)
		fmt.Println("===Start Run runClient2===")
	case 5:
		//go connectServant12()
		fmt.Println("===Start Run connectServant12===")
	default:
		fmt.Println("run option is not 1, 2 ,3 ,4 ,5")
	}

	for {
		time.Sleep(10 * time.Second)
	}
}

func runServant1() {
	//servant1.CreateNode(0, 12136, "127.0.0.1")
	servant1.createServHost(0, 12136, "127.0.0.1")
	servant1.Listen()
}

func runServant2() {
	//servant2.CreateNode(0, 12137, "127.0.0.1")
	servant2.createServHost(0, 12137, "127.0.0.1")
	servant2.Listen()
}

func runClient1(addr string) {
	//client1.CreateNode(1, 0, "")
	client1.createServHost(1, 12136, "127.0.0.1")
	client1.ConnectTo(addr)
}
func runClient2(addr string) {
	client2.createServHost(1, 12137, "127.0.0.1")
	client2.ConnectTo(addr)
}
