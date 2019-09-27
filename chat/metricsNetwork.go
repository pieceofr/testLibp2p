package main

import (
	"fmt"
	"time"

	"github.com/bitmark-inc/bitmarkd/counter"
	"github.com/bitmark-inc/logger"
	p2pcore "github.com/libp2p/go-libp2p-core"
	p2pnet "github.com/libp2p/go-libp2p-core/network"
	multiaddr "github.com/multiformats/go-multiaddr"
)

// MetricsNetwork contain P2P network metrics
type MetricsNetwork struct {
	log         *logger.L
	streamCount counter.Counter
	connCount   counter.Counter
}

//NewMetricsNetwork return a metricNetwork
func NewMetricsNetwork() MetricsNetwork {
	return MetricsNetwork{}
}

func (m *MetricsNetwork) startMonitor(host p2pcore.Host) {
	host.Network().Notify(&p2pnet.NotifyBundle{
		ListenF: func(net p2pnet.Network, addr multiaddr.Multiaddr) {
			fmt.Printf("@@Host Listen: %v is listen at %v\n", addr.String(), time.Now())
		},
		ConnectedF: func(net p2pnet.Network, conn p2pnet.Conn) {
			m.connCount.Increment()
			fmt.Printf("@@: Conn Conn: %v Connected at %v ConnCount:%d\n", conn.RemoteMultiaddr().String(), time.Now(), m.connCount)
		},
		DisconnectedF: func(net p2pnet.Network, conn p2pnet.Conn) {
			m.connCount.Decrement()
			fmt.Printf("@@Conn Disconn: %v Disconnected at %v  ConnCount:%d\n", conn.RemoteMultiaddr().String(), time.Now(), m.connCount)
		},
		OpenedStreamF: func(net p2pnet.Network, stream p2pnet.Stream) {
			m.streamCount.Increment()
			fmt.Printf("@@Stream Opened: %v-%v is Opened at %v streamCount:%d\n", stream.Conn().RemoteMultiaddr().String(), stream.Protocol(), time.Now(), m.streamCount)
		},
		ClosedStreamF: func(net p2pnet.Network, stream p2pnet.Stream) {
			m.streamCount.Decrement()
			fmt.Printf("@@Stream Closed:%v-%v is Closed at %v streamCount:%d\n", stream.Conn().RemoteMultiaddr().String(), stream.Protocol(), time.Now(), m.streamCount)
		},
	})
}
