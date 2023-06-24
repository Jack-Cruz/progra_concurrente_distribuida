package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	con, _ := net.Dial("tcp", "localhost:8000")
	defer con.Close()
	r := bufio.NewReader(con)
	g := &Game{[]rune{'O', 'X'},
	           0,
			   [][]rune{{' ', ' ', ' '},
			            {' ', ' ', ' '},
			            {' ', ' ', ' '}}}
	var i, j int
	for {
		g.Print()
		for {
			fmt.Print("Jugada [fila columna]: ")
			fmt.Scanf("%d %d\n", &i, &j)
			if g.Move(i, j) {
				break
			}
		}
		g.Print()
		buf, _ := json.Marshal(*g)
		fmt.Println(string(buf))
		fmt.Fprintln(con, string(buf))
		msg, _ := r.ReadString('\n')
		_ = json.Unmarshal([]byte(msg), g)
		p := g.Check()
		if p != ' ' {
			fmt.Printf("Winner %c\n", p)
			break
		}
	}
}
