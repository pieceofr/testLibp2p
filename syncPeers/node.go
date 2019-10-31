package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	p2pcore "github.com/libp2p/go-libp2p-core"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/peer"
	peerstore "github.com/libp2p/go-libp2p-core/peerstore"
	tls "github.com/libp2p/go-libp2p-tls"
	multiaddr "github.com/multiformats/go-multiaddr"
)

var nodeProtocol = multiaddr.ProtocolWithCode(multiaddr.P_P2P).Name

//SimpleNode 0:servant 1:client
type SimpleNode struct {
	Host       p2pcore.Host
	NodeType   int
	ListenAddr string
	PublicIP   string
	Port       int
	MetricsNetwork
}

func (n *SimpleNode) createServHost(nodeType int, port int, ip string, keypath string) {
	n.PublicIP = ip
	n.Port = port
	prvKey, err := loadIdentity(keypath)
	if err != nil {
		fmt.Println("LoadIdentity Error:", err)
		prvKey, _, err = crypto.GenerateEd25519Key(rand.Reader)
		if err != nil {
			panic(err)
		}
	}

	// 0.0.0.0 will listen on any interface device.

	// libp2p.New constructs a new libp2p Host.
	// Other options can be added here.
	options := []libp2p.Option{libp2p.Identity(prvKey), libp2p.Security(tls.ID, tls.New)}
	if 0 == nodeType {
		sourceMultiAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
		if err != nil {
			fmt.Println(err)
		}
		options = append(options, libp2p.ListenAddrs(sourceMultiAddr))
	}
	host, err := libp2p.New(context.Background(),
		options...,
	)
	if err != nil {
		panic(err)
	}
	n.Host = host
	n.MetricsNetwork = NewMetricsNetwork()
	go n.startMonitor(n.Host)
}

//Listen Create a Listener
func (n *SimpleNode) Listen() {
	var shandler SimpleStream
	shandler.ID = fmt.Sprintf("%s", n.Host.ID())
	n.ListenAddr = fmt.Sprintf("/ip4/%s/tcp/%v/p2p/%s", n.PublicIP, n.Port, n.Host.ID().Pretty())
	fmt.Printf("A servant:%s is listen to \n%s\n", n.Host.ID(), n.ListenAddr)
	n.Host.SetStreamHandler("/chat/1.0.0", shandler.handleStream)

	for _, la := range n.Host.Network().ListenAddresses() {
		if _, err := la.ValueForProtocol(multiaddr.P_TCP); err == nil {
			fmt.Print(err)
			return
		}
		for _, a := range n.Host.Addrs() {
			fmt.Printf("\x1b[32mHost Address: %s/%v/%s\x1b[0m\n", a, nodeProtocol, n.Host.ID())
		}
		fmt.Printf("\nWaiting for incoming connection\n\n")
	}
	// Hang forever
	<-make(chan struct{})
}

//ConnectTo Connect to an address
func (n *SimpleNode) ConnectTo(addr string) {
	fmt.Println("Trying to Connect to ", addr)
	maddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		log.Fatalln()
		fmt.Println("ConnectTo Error:", err)
	}
	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		log.Fatalln(err)
		fmt.Println("ConnectTo Error AddrInfoFromP2pAddr:", err)
	}
	n.Host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
	s, err := n.Host.NewStream(context.Background(), info.ID, "/chat/1.0.0")
	if err != nil {
		panic(err)
	}
	// Create a thread to read and write data.
	shandler := SimpleStream{ID: n.Host.ID().String()}
	//shandler.ID = fmt.Sprintf("%s", n.Host.ID())
	shandler.registerStream(s)

}

func (n *SimpleNode) showHostAddr() {
	for _, la := range n.Host.Addrs() {
		fmt.Printf(" - %v\n", la)
	}
	fmt.Println()
}

func (n *SimpleNode) shortID() string {
	id := n.Host.ID()
	return fmt.Sprintf("%s", id[len(id)-10:len(id)-1])
}
