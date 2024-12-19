package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var debug bool
var minport, maxport int
var ip string
var log = logrus.New()

func init() {
	flag.IntVar(&minport, "l", 0, "Min port number")
	flag.IntVar(&maxport, "u", 4096, "Max port number")
	flag.BoolVar(&debug, "d", false, "Debug Log level")
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

	ports := []int{}
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}

	for i := minport; i < maxport; i++ {
		wg.Add(1)
		go ScanTCPPort(ip, i, &ports, mutex, wg)

	}

	elapsed := time.Since(start)
	fmt.Println("Scanned from ", minport, " to ", maxport, " in ", elapsed)
	fmt.Println(ports)
}
