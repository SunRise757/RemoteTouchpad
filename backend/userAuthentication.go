package main

import (
	"log"
	"net"
	"os"
	"strings"
)

func getLanIp() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		return ""
	}

	//	var ips []string
	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, _ := iface.Addrs()
		for _, addr := range addrs {
			ip, _, err := net.ParseCIDR(addr.String())
			if err == nil && ip.To4() != nil && strings.HasPrefix(ip.String(), "192.168.") {
				return ip.String()
				//ips = append(ips, ip.String())
			}
		}
	}
	return ""
}

