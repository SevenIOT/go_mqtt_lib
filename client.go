package main

import (
	"fmt"
	"os"
	"net"
	"time"
)


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1883")
	if  err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	defer conn.Close()
	conn.Write([]byte("Hello world!\n"))
	for ; ;  {
		content := "hello"
		_, err := conn.Write([]byte(content))
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			break
		}
		fmt.Println(content)
		time.Sleep(time.Second)
	}
}