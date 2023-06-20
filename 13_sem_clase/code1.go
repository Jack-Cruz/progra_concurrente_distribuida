package main

import (
	"flag"
	"fmt"
	"strings"
	"net"
)

func listen(hostname string, chRemotes chan[]string){
	if ln, err := net.Listen("tcp", hostname); err != nil {
		defer ln.Close()
		fmt.Println("Listening...")
		for {
			if cn, err := ln.Accept(); err != nil {
				go handle(cn, chRemotes)
			}
		}
	}
}

func handle(cn net.Conn, chRemotes chan[] string) {
	defer cn.Close()
	fmt.Println("Connection accepted from \n", cn.RemoteAddr())
	remotes := <- chRemotes
	fmt.Println(remotes)
}

func main() {
	rawhostname := flag.String("h", "localhost:8000", "tdb")
	rawremotes := flag.String("r", "", "tdb")
	flag.Parse()

	if *rawhostname == "" || *rawremotes == "" {
		flag.PrintDefaults()
		return
	}
	chRemotes := make(chan[] string)
	chRemotes <- strings.Split(strings.TrimSpace(*rawremotes), " ")
	hostname := *rawhostname

	listen(hostname, chRemotes)
}