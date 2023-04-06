package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 解析命令行参数
	startIP := flag.String("IP", "", "起始IP地址或地址段")
	outputFile := flag.String("O", "ip.txt", "输出文件名")
	flag.Parse()

	if *startIP == "" || *outputFile == "" {
		fmt.Println("\"使用方法: \\n Windows: Piplist.exe -IP <ip Interval> -O <out file> \\n Mac or Linux: Piplist -IP <ip Interval> -O <out file>\\n 更多高级用法请参考GitHub：https://github.com/Moxin1044/Piplist\\n MIT License Copyright (c) 2023 末心\"")
		os.Exit(1)
	}

	// 解析起始IP地址或地址段
	var ips []string
	if strings.Contains(*startIP, "-") {
		ipRange := strings.Split(*startIP, "-")
		if len(ipRange) != 2 {
			fmt.Println("Invalid IP range")
			os.Exit(1)
		}
		start := net.ParseIP(ipRange[0])
		if start == nil {
			fmt.Println("Invalid start IP address")
			os.Exit(1)
		}
		end := net.ParseIP(ipRange[1])
		if end == nil {
			fmt.Println("Invalid end IP address")
			os.Exit(1)
		}
		for ip := start; !ip.Equal(end); inc(ip) {
			ips = append(ips, ip.String())
		}
		ips = append(ips, end.String())
	} else {
		_, ipnet, err := net.ParseCIDR(*startIP)
		if err != nil {
			fmt.Println("Invalid CIDR format")
			os.Exit(1)
		}
		for ip := ipnet.IP.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			ips = append(ips, ip.String())
		}
	}

	// 写入IP地址列表到文件
	f, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Failed to create output file:", err)
		os.Exit(1)
	}
	defer f.Close()

	for _, ip := range ips {
		_, err := f.WriteString(ip + "\n")
		if err != nil {
			fmt.Println("Failed to write IP address to file:", err)
			os.Exit(1)
		}
	}
}

// 将IP地址加1
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
