package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func PrintPorts(ports []PortInfo) {
	fmt.Println("----------------------------------------------------------")
	for _, p := range ports {
		fmt.Print(p.port, " ")
	}
	fmt.Println("\n----------------------------------------------------------")
}

// PrintPortInfo 打印PortInfo类型的slice
func PrintPortInfo(ports []PortInfo) {
	fmt.Printf("%-10s %-20s %-15s\n", "Port", "Protocol", "IP Address")
	fmt.Println("----------------------------------------------------------")
	for _, p := range ports {
		fmt.Printf("%-10d %-20s %-15s\n", p.port, p.protocol, p.ip)
	}
}

func PrintPortTable(ports []PortInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Port", "Remote Address", "Protocol"})
	table.SetBorder(true)
	table.SetHeaderLine(true)
	// table.SetCenterSeparator("|")
	// table.SetColumnSeparator("|")
	// table.SetRowSeparator("-")
	// table.SetTablePadding(" ")

	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_LEFT,   // Right align numbers (Port)
		tablewriter.ALIGN_CENTER, // Center align text (Protocol)
		tablewriter.ALIGN_CENTER, // Center align text (Status)
	})

	for _, p := range ports {
		//处理字符串中间的空格
		displayProtocol := strings.ReplaceAll(p.protocol, "\n", " ")
		trimmedStr := strings.TrimSpace(displayProtocol)
		str := truncateString(trimmedStr, 20)

		table.Append([]string{fmt.Sprintf("%d", p.port), p.ip, str})
	}

	table.Render()

}

// truncateString 截断字符串到指定长度
// Table包有问题 无法处理过长的字符串 20是一个比较靠谱的长度
func truncateString(str string, maxLength int) string {
	if len(str) > maxLength {
		return str[:maxLength] + "..."
	}
	return str
}
