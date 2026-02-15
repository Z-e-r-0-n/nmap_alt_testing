package main

import (
	"fmt"
	"time"
)

func main() {
	myphone := make(chan string)
	for i := range 10000 {
		time.Sleep(2 * time.Millisecond)
		go Sccan("192.168.14.2:", fmt.Sprint(i), myphone)
	}
	for range 10000 {
		<-myphone
	}

	fmt.Println("Hello World!")
}

