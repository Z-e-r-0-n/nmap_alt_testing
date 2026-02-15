package main

import (
	"fmt"
	"net"
	"time"
)

func Sccan(ip string, port string, phone chan string) {
	times := time.Second * 6
	conn, err := net.DialTimeout("tcp", ip+port, times)
	if err != nil {

		phone <- ""
		return
	}
	conn.SetReadDeadline(time.Now().Add(2 * time.Second))

	buffer := make([]byte, 1024)
	
	n, err := conn.Read(buffer)
	conn.Close()
	phone <- ""
	fmt.Println(port + "open" +  string(buffer[:n]))
	

	

}

