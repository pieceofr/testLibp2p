package main

import (
	"context"
	"fmt"

	"crypto/rand"

	libp2p "github.com/libp2p/go-libp2p"
	p2pcore "github.com/libp2p/go-libp2p-core"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peerlib "github.com/libp2p/go-libp2p-core/peer"
	peerstore "github.com/libp2p/go-libp2p-core/peerstore"
	protocol "github.com/libp2p/go-libp2p-protocol"
	tls "github.com/libp2p/go-libp2p-tls"
	"github.com/multiformats/go-multiaddr"
)

var servHost p2pcore.Host
var nodeProtocol = "p2p"
var addrs = []string{
	"/ip4/127.0.0.1/tcp/2136/ipfs/12D3KooWCvzWT8L6HTtAwac5Vr7RzLitbyZkJ9T9fnikWonnS1DR",
	"/ip4/127.0.0.1/tcp/12136/p2p/12D3KooWCvzWT8L6HTtAwac5Vr7RzLitbyZkJ9T9fnikWonnS1DR",
}

func main() {
	servHost = *newClientHost()
	connectTo(addrs[1])
}

func newClientHost() *p2pcore.Host {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}
	options := []libp2p.Option{libp2p.Identity(priv), libp2p.Security(tls.ID, tls.New)}
	newNode, err := libp2p.New(ctx, options...)
	if err != nil {
		fmt.Println(err)
	}
	return &newNode
}

func connectTo(addr string) {
	fmt.Println("Trying to Connect to ", addr)
	maddr, err := multiaddr.NewMultiaddr(addr)
	if err != nil {
		fmt.Println(err)
	}
	info, err := peerlib.AddrInfoFromP2pAddr(maddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	servHost.Peerstore().AddAddrs(info.ID, info.Addrs, peerstore.PermanentAddrTTL)

	// Start a stream with the destination.
	// Multiaddress of the destination peer is fetched from the peerstore using 'peerId'.
	s, err := servHost.NewStream(context.Background(), info.ID, protocol.ID(nodeProtocol))
	if err != nil {
		panic(err)
	}
	// Create a thread to read and write data.
	var shandler SimpleStream

	shandler.ID = fmt.Sprintf("%s", servHost.ID())
	shandler.handleStream(s)
}
