package main

import (
	"fmt"
	"net"
	"sync"
)

func ScanTCPPort(ip string, port int, ports *[]int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	log.Debug("Connecting: ", address)

	conn, err := net.Dial("tcp", address)
	// Don't use DialTimeout
	if err == nil {
		defer conn.Close()
		mutex.Lock()
		*ports = append(*ports, port)
		mutex.Unlock()
	}

}
