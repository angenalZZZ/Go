package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/pion/webrtc"
)

func main() {
	signalAddr := flag.String("signal", "hare1039.nctu.me:6666", "Signal server")
	stunAddr := flag.String("stun", "hare1039.nctu.me:3478", "STUN server")
	offer := flag.Bool("offer", false, "connect to exposed service")
	exposeAddr := flag.String("expose", "localhost:22", "exposed service")
	listenerAddr := flag.String("listen", ":10000", "local listener for remote service(e.g. ssh)")
	help := flag.Bool("help", false, "help")
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *offer {
		fmt.Println("offer (client) mode")
	} else {
		fmt.Println("answer (server) mode")
	}

	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:" + *stunAddr, "stun:stun.l.google.com:19302"},
			},
		},
	}

	if !*offer {
		// serve peer candidate
		peer := exposeServer(config, *exposeAddr)
		offerChan, answerChan := connectSignal(*signalAddr)
		offer := <-offerChan

		if err := peer.SetRemoteDescription(offer); err != nil {
			panic(err)
		}

		// Create answer
		answer, err := peer.CreateAnswer(nil)
		if err != nil {
			panic(err)
		}
		if err = peer.SetLocalDescription(answer); err != nil {
			panic(err)
		}

		// Send the answer
		answerChan <- answer
	} else {
		peer := offerClient(config, *listenerAddr)
		// Create an offer to send to the browser
		offer, err := peer.CreateOffer(nil)
		if err != nil {
			panic(err)
		}

		// Sets the LocalDescription, and starts our UDP listeners
		err = peer.SetLocalDescription(offer)
		if err != nil {
			panic(err)
		}

		// Exchange the offer for the answer
		answer := offerSignal(offer, *signalAddr)

		// Apply the answer as the remote description
		err = peer.SetRemoteDescription(answer)
		if err != nil {
			panic(err)
		}
	}
	// Block forever
	select {}
}

func exposeServer(config webrtc.Configuration, exposeAddr string) *webrtc.PeerConnection {
	conn, err := net.Dial("tcp", exposeAddr)

	peer, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	peer.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State changed: %s\n", state.String())
	})

	peer.OnDataChannel(func(d *webrtc.DataChannel) {
		fmt.Printf("New DataChannel %s %d\n", d.Label(), d.ID())

		d.OnOpen(func() {
			fmt.Printf("Data channel '%s'-'%d' open", d.Label(), d.ID())

			buf := make([]byte, 8192)
			reader := bufio.NewReader(conn)
			for {
				n, err := reader.Read(buf)
				if err != nil {
					panic(err)
				}

				slice := buf[0:n]
				d.Send(slice)
			}
		})

		d.OnMessage(func(msg webrtc.DataChannelMessage) {
			conn.Write(msg.Data)
		})
	})

	return peer
}

func connectSignal(signalServer string) (offerOut chan webrtc.SessionDescription, answerIn chan webrtc.SessionDescription) {
	offerOut = make(chan webrtc.SessionDescription)
	answerIn = make(chan webrtc.SessionDescription)

	go func() {
		conn, err := net.Dial("tcp", signalServer)
		if err != nil {
			panic(err)
		}

		var offer webrtc.SessionDescription
		err = json.NewDecoder(conn).Decode(&offer)
		if err != nil {
			panic(err)
		}

		offerOut <- offer
		answer := <-answerIn
		err = json.NewEncoder(conn).Encode(answer)
		if err != nil {
			panic(err)
		}

		conn.Write([]byte("\n"))
	}()

	return
}

func offerClient(config webrtc.Configuration, listenerAddr string) *webrtc.PeerConnection {
	fmt.Println("Listen on", listenerAddr)
	ln, err := net.Listen("tcp", listenerAddr)
	if err != nil {
		panic(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	bufio.NewReader(conn)

	peer, err := webrtc.NewPeerConnection(config)
	if err != nil {
		panic(err)
	}

	// label "data"
	d, err := peer.CreateDataChannel("data", nil)
	if err != nil {
		panic(err)
	}

	peer.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		fmt.Printf("ICE Connection State changed: %s\n", state.String())
	})

	d.OnOpen(func() {
		fmt.Printf("Data channel '%s'-'%d' open.", d.Label(), d.ID())
		buf := make([]byte, 8192)
		reader := bufio.NewReader(conn)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				panic(err)
			}

			slice := buf[0:n]
			d.Send(slice)
		}
	})

	d.OnMessage(func(msg webrtc.DataChannelMessage) {
		conn.Write(msg.Data)
	})

	return peer
}

func offerSignal(offer webrtc.SessionDescription, signalServer string) webrtc.SessionDescription {
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(offer)
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", signalServer)
	if err != nil {
		panic(err)
	}

	conn.Write(b.Bytes())
	conn.Write([]byte("\n"))

	var answer webrtc.SessionDescription
	err = json.NewDecoder(conn).Decode(&answer)
	if err != nil {
		panic(err)
	}

	return answer
}
