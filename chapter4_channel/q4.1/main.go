package main

import (
	"fmt"
	"time"
)

const (
	ten = 10
)
func main() {
	timer := time.After(ten * time.Second)
	fmt.Println("Start")
	<-timer
	fmt.Println("It's 10 seconds")
}
