package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func descubrirIP() {
	ifaces, err := net.Interfaces()

	if err != nil {
		log.Print(fmt.Errorf("Error de dirección Red: %v\n", err.Error()))
		return
	}

	for _, i := range ifaces {
		dirs, err := i.Addrs()
		if err != nil {
			log.Print(fmt.Errorf("Error de dirección Red: %v\n", err.Error()))
			continue
		}
		for _, dir := range dirs {
			fmt.Printf("%v : %v\n", i.Name, dir)
		}
	}
}

func descubrirIP2() string {
	ifaces, _ := net.Interfaces()

	for _, iface := range ifaces {
		//if strings.HasPrefix(iface.Name, "eth0") {
		if strings.HasPrefix(iface.Name, "Ethernet") {
			dirs, _ := iface.Addrs()
			for _, dir := range dirs {
				// fmt.Printf("%v : %v\n", iface.Name, dir)
				switch t := dir.(type) {
				case *net.IPNet:
					if t.IP.To4() != nil {
						// fmt.Printf("%v : %v\n", iface.Name, t.IP.To4().String())
						return string(t.IP.To4().String())
					}
					break
				case *net.IPAddr:
					fmt.Printf("%v : %v\n", iface.Name, t.IP)
					break
				}
			}
		}
	}
	return ""
}

func main() {
	// descubrirIP()
	fmt.Println(descubrirIP2())
}
