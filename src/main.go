package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var debug bool
var minport, maxport, dt int
var ip string
var log = logrus.New()

func init() {
	flag.IntVar(&minport, "l", 0, "Min port number")
	flag.IntVar(&maxport, "u", 65536, "Max port number")
	flag.IntVar(&dt, "dt", 0, "Port Access interval in ms, -1 for no wait")
	flag.BoolVar(&debug, "debug", false, "Debug Log level")
	flag.StringVar(&ip, "ip", "127.0.0.1", "IP address")
	flag.Parse()

	log.SetLevel(logrus.InfoLevel)
	if debug {
		log.SetLevel(logrus.DebugLevel)
	}

	formatter := &logrus.TextFormatter{
		ForceColors: true,
	}
	log.SetFormatter(formatter)
	log.Info("Program Start")
}

func main() {
	start := time.Now()

	ports := []PortInfo{}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	var ticker *time.Ticker
	ticker = nil
	if dt > 0 {
		log.Info("Using wait interval")
		ticker = time.NewTicker(time.Duration(dt) * time.Millisecond)
		defer ticker.Stop()
	}

	for i := minport; i < maxport; i++ {
		wg.Add(1)
		if ticker == nil {
			go ScanTCPPort(ip, i, &ports, mutex, wg)
		} else {
			go ScanTCPPortDT(ip, i, &ports, mutex, wg, ticker)
		}
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Println("Scanned from ", minport, " to ", maxport, " in ", elapsed)
	fmt.Println(ports)
}
