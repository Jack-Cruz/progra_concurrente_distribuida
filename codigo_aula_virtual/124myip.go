package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println(myIp())
}
func myIp() string { // mandrakeando ando
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		print(iface.Name, "\n")
		//if strings.HasPrefix(iface.Name, "Local") {
		//if strings.HasPrefix(iface.Name, "Wi-Fi") {
		if strings.HasPrefix(iface.Name, "Ethernet") {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					return v.IP.String()
				case *net.IPAddr:
					return v.IP.String()
				}
			}
		}
	}
	return ""
}