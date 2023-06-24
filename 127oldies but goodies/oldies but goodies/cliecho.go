package main

import (
    "fmt"
    "net"
    "bufio"
    "os"
)

func main() {
    con, _ := net.Dial("tcp", "localhost:8000")
    defer con.Close()
    r := bufio.NewReader(con)
    gin := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Msg: ")
        msg, _ := gin.ReadString('\n')
        fmt.Fprint(con, msg)
        resp, _ := r.ReadString('\n')
        fmt.Print("Resp:", resp)
        if len(msg) == 0 || msg[0] == 'x' {
            break
        }
    }
}
