package main

import (
	"bytes"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/go-ping/ping"
)

// PingOk Ping命令模式
func PingWithCommand(host string) bool {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("ping", "-c", "1", "-W", "1", host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		log.Info(out.String())
		if strings.Contains(out.String(), "ttl=") {
			return true
		}
	case "windows":
		// 设置控制台编码为UTF-8
		cmd := exec.Command("cmd", "/C", "chcp 65001")
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Run()
		cmd = exec.Command("ping", "-n", "1", "-w", "500", host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()

		// 将输出转换为UTF-8编码
		output := out.String()
		if !utf8.ValidString(output) {
			output = string([]rune(output))
		}

		log.Info(output)
		if strings.Contains(out.String(), "TTL=") {
			return true
		}
	case "darwin":
		cmd := exec.Command("ping", "-c", "1", "-t", "1", host)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Run()
		log.Info(out.String())
		if strings.Contains(out.String(), "ttl=") {
			return true
		}
	}
	return false
}

// 直接发ICMP包 使用ICMP库
func PingWithIcmp(host string) bool {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return false
	}
	pinger.SetPrivileged(true)
	pinger.Count = 1
	pinger.Timeout = 800 * time.Millisecond
	if pinger.Run() != nil { // Blocks until finished. return err
		return false
	}
	if stats := pinger.Statistics(); stats.PacketsRecv > 0 {
		return true
	}
	return false
}
