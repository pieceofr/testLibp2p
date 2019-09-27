package main

import (
	"bufio"
	"fmt"
	"log"
	"time"

	"github.com/libp2p/go-libp2p-core/network"
)

//SimpleStream for Adding more infomation for handleStream
type SimpleStream struct {
	ID string
}

func (s *SimpleStream) handleStream(stream network.Stream) {
	log.Println("Start a Listening stream!")

	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	for {
		str, _ := rw.ReadString('\n')
		if str == "" {
			return
		}
		if str != "\n" {
			fmt.Printf("%s RECIEVE:\x1b[32m%s\x1b[0m> ", s.ID, str)
		}
	}
	// stream 's' will stay open until you close it (or the other side closes it).
}

func (s *SimpleStream) registerStream(stream network.Stream) {
	log.Println("Start a Registering stream!")
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	rw.WriteString(fmt.Sprintf("%s SEND %s\n", s.ID, time.Now().UTC().String()))
	rw.Flush()
}

func readData(rw *bufio.ReadWriter, who string) {
	for {
		str, _ := rw.ReadString('\n')

		if str == "" {
			return
		}
		if str != "\n" {
			fmt.Printf("%s RECIEVE:\x1b[32m%s\x1b[0m> ", who, str)
		}
	}
}

func writeData(rw *bufio.ReadWriter, who string) {
	timeUp := time.After(5 * time.Second)
	for {
		select {
		case <-timeUp:
			rw.WriteString(fmt.Sprintf("%s SEND %s\n", who, time.Now().UTC().String()))
			rw.Flush()
		}
	}
}
