package main

import (
	"fmt"
	"net"
	"bufio"
	"strconv"
	//"time"
)

const MAX = 5

func server(ch, end chan int, addr string) {
	ln, _ := net.Listen("tcp", addr)
	defer ln.Close()
	con, _ := ln.Accept()
	defer con.Close()
	r := bufio.NewReader(con)
	var num int
	for {
		strNum, _ := r.ReadString('\n')
		fmt.Println("Str: ", strNum)
		if (len(strNum) == 0) {
			num = -1
		} else {
			num, _ = strconv.Atoi(strNum[:len(strNum) - 1])
		}
		fmt.Println(num)
		if num == 0 {
			fmt.Println("Ouch, I lost! BOOOM!")
		} else {
			//time.Sleep(time.Second)
		}
		ch<- num
		if num < 0 {
			break
		}
	}
	end<- 0
}
func client(ch chan int, addr string) {
	var con net.Conn
	created := false
	for {
		num := <-ch
		if !created {
			created = true
			con, _ = net.Dial("tcp", addr)
			defer fmt.Println("Cerrando conexion a NEXT Nodo")
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
	fmt.Print("Ingrese un numero: ")
	fmt.Scanf("%d\n", &num)
	ch<- num
}

func main() {
	end := make(chan int)
	ch := make(chan int)
	go server(ch, end, "10.11.98.229:8000")
	go client(ch, "10.11.98.205:8000")
	go start(ch)
	<-end
}