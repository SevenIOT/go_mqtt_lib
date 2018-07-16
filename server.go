package main

import (
	"net"
	"fmt"
	"strconv"
)


func handleConn(conn net.Conn) {
	//  var buf [50]byte
	buf := make([]byte, 1024)
	defer conn.Close()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("conn closed: %s", err.Error())
			return
		}
		fmt.Println("recv msg: ", string(buf[0:n]))
		writeBuf := []byte {'a', 'b'}
		conn.Write(writeBuf)
	}
}

func main() {
	port := 1883
	listen, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	fmt.Printf("server listen port: %d\n", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		go handleConn(conn)
	}
}