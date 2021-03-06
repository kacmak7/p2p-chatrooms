package main

import (
	"log"
	"os"
	"github.com/akamensky/argparse"
	"github.com/kacmak7/p2p-chatrooms/server"
	"github.com/kacmak7/p2p-chatrooms/client"
)

func main() {
	// Commands parser
	parser := argparse.NewParser("commands", "Frames commands")

	connectCmd := parser.NewCommand("connect", "Connect to a room")
	roomFlag := connectCmd.String("r", "room", &argparse.Options{Required: true, Help: "Room hash"})
	peerFlag := connectCmd.String("peer", "mate", &argparse.Options{Required: true, Help: "Peer to which you want to connect"})
	portFlag := connectCmd.String("p", "port", &argparse.Options{Required: true, Help: "Port for a connection to room. You can use a single port for multiple rooms"})

	sendCmd := parser.NewCommand("send", "Send a message")
	messageFlag := sendCmd.String("m", "message", &argparse.Options{Required: true, Help: "Message text"})
	roomFlag = sendCmd.String("r", "room", &argparse.Options{Required: true, Help: "Room hash"})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	if connectCmd.Happened() {
		server.Connect(*roomFlag, *portFlag)
	} else if sendCmd.Happened() {
		client.Send()
	}



}
