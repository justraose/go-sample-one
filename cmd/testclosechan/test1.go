package main

import (
	"fmt"
	"time"
)

func main() {
	ints := make(chan int)
	fmt.Println("1")
	go func(){
		time.Sleep(1000*1000*1000*10)
		close(ints)
	}()
	<-ints
	fmt.Println("2")
}
