package server

import (
	"log"
	"time"
	"net"
	"github.com/patrickmn/go-cache"

	"github.com/kacmak7/p2p-chatrooms/utils"
)

func handle(c net.Conn, roomHash string, port string) {
	log.Printf("Connecting to the room %s", roomHash)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	utils.CheckError("Not able to resolve TCP Address", err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckError("Failed to start a listener", err)
