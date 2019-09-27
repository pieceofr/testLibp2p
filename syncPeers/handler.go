package main

import (
	"bufio"
	"fmt"
	"io"
	"log"

	proto "github.com/golang/protobuf/proto"
	streammux "github.com/ipsn/go-ipfs/gxlibs/github.com/libp2p/go-stream-muxer"
	"github.com/libp2p/go-libp2p-core/network"
)

const (
	RED   = "\x1b[31m"
	GREEN = "\x1b[32m"
	CLEAR = "\x1b[0m"
)

//SimpleStream for Adding more infomation for handleStream
type SimpleStream struct {
	ID string
}

func (s *SimpleStream) handleStream(stream network.Stream) {
	log.Println("Got a new stream!")
	// Create a buffer stream for non blocking read and write.
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
loop:
	for {
		req := make([]byte, 1000)
		reqLen, err := rw.Read(req)
		if err != nil {
			if err == io.EOF {
				continue loop
			}
			if err == streammux.ErrReset {
				fmt.Println("Stream reset", err)
				break
			}
			fmt.Println("Stream Read Error:", err)
			break
		}

		fmt.Println("reqLen Lenght", reqLen)
		if reqLen == 0 {
			fmt.Println("Stream Len = 0")
		}
		reqMsg := MockProtoMessage{}
		err = proto.Unmarshal(req[:reqLen], &reqMsg)
		if err != nil {
			fmt.Println("\x1b[31mUnmarshal Error:\x1b[0m", err)
			continue loop
		}
		fmt.Printf("%s RECIEVE:\x1b[32m%s\x1b[0m >\n", s.ID[20:], reqMsg.Command)
		respMsg := MockProtoMessage{Command: "HiHi", Data: []byte("AAAsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssBB")}
		respPacked, err := proto.Marshal(&respMsg)
		if err != nil {
			fmt.Println("\x1b[31m Response Marshal Error:\x1b[0m", err)
			continue loop
		}
		_, err = rw.Write(respPacked)
		if err != nil {
			fmt.Println("Stream Read Error:", err)
		}
		rw.Flush()
	}
	// stream 's' will stay open until you close it (or the other side closes it).
}
func (s *SimpleStream) registerStream(stream network.Stream) {

	log.Println("Start a Registering stream!")
	rw := bufio.NewReadWriter(bufio.NewReader(stream), bufio.NewWriter(stream))
	req := MockProtoMessage{Command: "Hello", Data: []byte("I am register")}
	reqPacked, _ := proto.Marshal(&req)
	fmt.Println("reqPacked Lenght", len(reqPacked))
	//_, err := rw.Write(append(reqPacked))
	_, err := rw.Write(reqPacked)
	if err != nil {
		fmt.Println("Stream Read Error:", err)
	}
	rw.Flush()

	answer := make([]byte, 1000)
	ansLen, err := rw.Read(answer)
	if err != nil {
		if err == io.EOF {
			return
		}
		if err == streammux.ErrReset {
			fmt.Println("Stream reset", err)
			return
		}
		fmt.Println("Stream Read Error:", err)
		return
	}

	fmt.Println("ansLen Length", ansLen)
	if ansLen == 0 {
		fmt.Println("Stream Len = 0")
	}
	ansMsg := MockProtoMessage{}
	err = proto.Unmarshal(answer[:ansLen], &ansMsg)
	if err != nil {
		fmt.Println("\x1b[31mUnmarshal Error:\x1b[0m", err)
		return
	}
	fmt.Printf("%s RECIEVE:\x1b[32m%s\x1b[0m >\n", s.ID[20:], ansMsg.Command)
	defer func() {
		if err != nil {
			fmt.Println("\x1b[31mStream Closed\x1b[0m")
			//stream.Close() // Close 1 Side
			stream.Reset() // Close both side
		}
		fmt.Println("\x1b[32mStream Registered\x1b[0m")
	}()
}
