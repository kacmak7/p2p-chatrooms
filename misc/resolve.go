package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func main() {

    service := ":8586"
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)


	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		numRead, recvAddr, err := conn.ReadFrom(buf)
        	if err != nil {
            		fmt.Println(err)
		}
		if recvAddr != nil {
            		fmt.Println(recvAddr)
        	}
        	s := string(buf[:numRead])
              	fmt.Println(s)
	}
}

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
