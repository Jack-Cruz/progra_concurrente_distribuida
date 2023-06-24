package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", "localhost:8000")
	defer ln.Close()
	con, _ := ln.Accept()
	defer con.Close()
	r := bufio.NewReader(con)
	g := &Game{}
	var i, j int
	for {
		msg, _ := r.ReadString('\n')
		fmt.Print(msg)
		_ = json.Unmarshal([]byte(msg), g)
		p := g.Check()
		g.Print()
		if p != ' ' {
			fmt.Printf("Winner %c\n", p)
			break
		}
		for {
			fmt.Print("Jugada [fila columna]: ")
			fmt.Scanf("%d %d\n", &i, &j)
			if g.Move(i, j) {
				break
			}
		}
		g.Print()
		buf, _ := json.Marshal(*g)
		fmt.Fprintln(con, string(buf))
	}
}
