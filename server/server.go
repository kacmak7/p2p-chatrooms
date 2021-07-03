package server

import (
	"net"
	"log"
	//"github.com/patrickmn/go-cache"

	"github.com/kacmak7/p2p-chatrooms/utils"
)

func Connect(roomHash string, port string) {
	log.Printf("Connecting to the room %s", roomHash)

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":"+port)
	utils.CheckError("Not able to resolve TCP Address", err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	utils.CheckError("Failed to start a listener", err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		conn.Write([]byte("chatrooms here YOO!"))
		conn.Close()


	}
}
