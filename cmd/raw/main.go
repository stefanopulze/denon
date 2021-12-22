package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	denonIp := "192.168.1.22:23"

	c, err := net.Dial("tcp", denonIp)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	log.Println("Connected")
	_, err = c.Write([]byte("PWSTANDBY"))
	if err != nil {
		panic(err)
	}

	reply := make([]byte, 1024)

	_, err = c.Read(reply)
	if err != nil {
		panic(err)
	}

	fmt.Println("reply:", string(reply))
}
