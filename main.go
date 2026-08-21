package main

import (
	"fmt"
	"service"
)

func main() {
	fmt.Println("Hello, Bounty Hunter!")
	
	// Simulate concurrent registration attempts
	go service.RegisterService("service1")
	go service.RegisterService("service2")
	
	// Wait for registration to complete
	time.Sleep(2 * time.Second)
}
