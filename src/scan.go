package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// 保存端口信息的基础数据类型
type PortInfo struct {
	port     int
	protocol string
}

func ScanTCPPort(ip string, port int, ports *[]PortInfo, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)
	log.Debug("Connecting: ", address)

	conn, err := net.Dial("tcp", address)
	// Don't use DialTimeout
	if err == nil {
		defer conn.Close()
		result := PortInfo{port: port, protocol: readConnection(conn)}
		mutex.Lock()
		*ports = append(*ports, result)
		mutex.Unlock()

	}

}

// 使用Ticker降低执行时间的密度慢速执行
func ScanTCPPortDT(ip string, port int, ports *[]PortInfo, mutex *sync.Mutex, wg *sync.WaitGroup, ticker *time.Ticker) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", ip, port)

	if ticker == nil {
		log.Fatal("Error")
	}

	// 使用Ticker阻塞程序执行
	t := <-ticker.C
	log.Info(t, address)

	conn, err := net.Dial("tcp", address)
	// Don't use DialTimeout
	if err == nil {
		defer conn.Close()
		result := PortInfo{port: port, protocol: readConnection(conn)}
		mutex.Lock()
		*ports = append(*ports, result)
		mutex.Unlock()
	}

}

func readConnection(conn net.Conn) string {
	// 发送Ctrl+C表示结束会话
	_, err := conn.Write([]byte{'\x03'})
	if err != nil {
		log.Warn("Error sending Ctrl+C: ", err)
		return fmt.Sprint(err)
	}
	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // 设置读取超时
	n, err := conn.Read(buffer)
	if err == nil {
		log.Infof("%s Received data: %s ", conn.RemoteAddr().String(), string(buffer[:n]))
		return string(buffer[:n])

	} else {
		log.Warn("Error reading data: ", err)
		return fmt.Sprint(err)
	}
}
