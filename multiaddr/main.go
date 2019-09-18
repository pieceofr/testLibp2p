package main

import (
	"fmt"

	peerlib "github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
)

func main() {
	testFromRandomPrvKey()
	//testIDFromPrvKey()
	//testIDFromString(addrs)
}
func mockValidP2PAddress() []string {
	return []string{
		"/ip4/127.0.0.1/tcp/12146/p2p/12D3KooWCvzWT8L6HTtAwac5Vr7RzLitbyZkJ9T9fnikWonnS1DR",
		"/ip6/::1/tcp/12146/ipfs/QmR2ykuCBPY27hLUCJDa6KxVForfLUvz6AcCj33y5Hxx7D",
	}
}

func mockPrivateKey() []byte {
	privatekey := []byte("080112406eb84a3845d33c2a389d7fbea425cbf882047a2ab13084562f06875db47b5fdc2e45a298e6cd0472eeb97cd023c723824e157869d81039794864987c05b212a8")
	return privatekey
}
func testFromRandomPrvKey() {
	prvKey, _ := GenRandPrvKey()

	id, err := peerlib.IDFromPrivateKey(prvKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	idString := id.String() // hash ID base58Encoded
	idfrom58, err := peerlib.IDB58Decode(idString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("id from 58string:", idfrom58.String())
}

func testFromFilePrvKey() {
	prvBinary := mockPrivateKey()
	pkey, err := DecodeHexToPrvKey(prvBinary)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := peerlib.IDFromPrivateKey(pkey)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id.String())

}
func testIDFromString() {
	addrs := mockValidP2PAddress()
	var maAddrs []ma.Multiaddr
	for _, addr := range addrs {
		//		fmt.Println("index=", index, " addr=", addr)
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			fmt.Println("NewMultiaddr Error:", err)
		}
		maAddrs = append(maAddrs, maddr)
	}
	for _, maAddr := range maAddrs {
		info, err := peerlib.AddrInfoFromP2pAddr(maAddr)
		if err != nil {
			fmt.Println("IDFromString Error:", err)
		}
		fmt.Println("peerlib.ID.String:", info.ID)
	}

}

func testInfo(addrs []string) {
	for _, addr := range addrs {
		//		fmt.Println("index=", index, " addr=", addr)
		maddr, err := ma.NewMultiaddr(addr)
		if err != nil {
			panic(err)
		}
		fmt.Println("MultiAddr=", maddr.String())

		for i, p := range maddr.Protocols() {
			fmt.Println("protocol[", i, "]  code= ", p.Code, " name=", p.Name, " path=", p.Path)
			pv, err := maddr.ValueForProtocol(p.Code)
			if err != nil {
				fmt.Println("valueforprotocol error=", err)
				panic(err)
			}
			fmt.Println("pv:", pv)
		}

		info, err := peerlib.AddrInfoFromP2pAddr(maddr)
		if err != nil {
			fmt.Println("parse error=", err)
			panic(err)
		}
		fmt.Println("info.String()=", info.String())
		fmt.Println("info.ID", info.ID)
		fmt.Printf("info.Addrs=%v\n", info.Addrs)
	}

}
