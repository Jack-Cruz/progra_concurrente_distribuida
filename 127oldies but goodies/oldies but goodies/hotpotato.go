package main

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
	"time"
)

const MAX = 5

func server(ch, end chan int, port int) {
	ln, _ := net.Listen("tcp", fmt.Sprintf("localhost:800%d", port))
	defer ln.Close()
	con, _ := ln.Accept()
	defer con.Close()
	r := bufio.NewReader(con)
	for {
		strNum, _ := r.ReadString('\n')
		num, _ := strconv.Atoi(strNum[:len(strNum) - 1])
		fmt.Println(num)
		if num == 0 {
			fmt.Println("Ouch, I lost! BOOOM!")
		} else {
			time.Sleep(time.Second)
		}
		ch<- num
		if num < 0 {
			break
		}
	}
	end<- 0
}
func client(ch chan int, port int) {
	var con net.Conn
	created := false
	for {
		num := <-ch
		if !created {
			created = true
			con, _ = net.Dial("tcp", fmt.Sprintf("localhost:800%d", port))
			defer con.Close()
		}
		fmt.Fprintf(con, "%d\n", num - 1)
		if num < 0 {
			fmt.Println("UF!")
			break
		}
	}
}
func start(ch chan int) {
	var num int
	fmt.Print("Ingrese un número: ")
	fmt.Scanf("%d\n", &num)
	ch<- num
}

func main() {
	var port int
	fmt.Print("Puerto: ")
	fmt.Scanf("%d\n", &port)
	end := make(chan int)
	ch := make(chan int)
	go server(ch, end, port)
	go client(ch, (port + 1) % MAX)
	go start(ch)
	<-end
}