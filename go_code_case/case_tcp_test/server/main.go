package main

import (
	"fmt"
	"net"
)

func main() {
	// listen local port that start the service
	listener, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("start tcp server on 127.0.0.1:2000 faild,err: ", err)
	}

	// for listen connection from client
	for {
		// wait for connection from other client
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accpet faild,err:", err)
			return
		}
		// goroutine for connection
		go processConn(conn)
	}

}

func processConn(conn net.Conn) {
	// talk to client
	var tmp [128]byte

	for {
		n, err := conn.Read(tmp[:])
		if err != nil {
			fmt.Println("read from conn faild,err", err)
			return
		}
		fmt.Println(string(tmp[:n]))
	}
}
