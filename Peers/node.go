package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	p2pcore "github.com/libp2p/go-libp2p-core"
	"github.com/libp2p/go-libp2p-core/peer"
	crypto "github.com/libp2p/go-libp2p-crypto"
	peerstore "github.com/libp2p/go-libp2p-peerstore"
	tls "github.com/libp2p/go-libp2p-tls"
	"github.com/multiformats/go-multiaddr"
)

//SimpleNode 0:servant 1:client
type SimpleNode struct {
	Host       p2pcore.Host
	NodeType   int
	ListenAddr string
	PublicIP   string
	Port       int
}

func (n *SimpleNode) createServHost(nodeType int, port int, ip string) {
	n.PublicIP = ip
	n.Port = port
	prvKey, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
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
		fmt.Printf("Run './Peers -run=  -addr=/ip4/%s/tcp/%v/p2p/%s' on another console.\n", n.PublicIP, n.Port, n.Host.ID().Pretty())
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
		log.Fatalln(err)
	}

	info, err := peer.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		log.Fatalln(err)
	}
	n.Host.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
	s, err := n.Host.NewStream(context.Background(), info.ID, "/chat/1.0.0")
	if err != nil {
		panic(err)
	}
	// Create a thread to read and write data.
	var shandler SimpleStream

	shandler.ID = fmt.Sprintf("%s", n.Host.ID())
	shandler.handleStream(s)

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
