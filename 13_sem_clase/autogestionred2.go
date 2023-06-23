package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

// bitacora de direcciones de red (miembros de la red)
var addrs []string

var hostaddr string

const (
	registerport = 8000
	notifyport   = 8001
)

// Obtener mi ip
func getMyIP() string {
	ifaces, _ := net.Interfaces()

	for _, iface := range ifaces {
		if strings.HasPrefix(iface.Name, "eth0") {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				switch t := addr.(type) {
				case *net.IPNet:
					return t.IP.String()
				case *net.IPAddr:
					return t.IP.String()
				}
			}
		}
	}
	return ""
}

// Servidor de registro
func registerServer() {
	hostname := fmt.Sprintf("%s:%d", hostaddr, registerport)
	ln, _ := net.Listen("tcp", hostname)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleRegister(conn)
	}
}

func handleRegister(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	remoteIP, _ := br.ReadString('\n')
	remoteIP = strings.TrimSpace(remoteIP)
	fmt.Printf("Nuevo host en la red: %s\n", remoteIP)
	
	// Enviar mi bitacora de addrs al nuevo host
	addrbytes, _ := json.Marshal(addrs)
	fmt.Fprintf(conn, "%s\n", string(addrbytes))

	// Notificar a todos que llego un nuevo host
	tellEverybody(remoteIP)

	// Agregar el nuevo host a mi bitacora
	addrs = append(addrs, remoteIP)
}

func tellEverybody(remoteIP string) {
	for _, addr := range addrs {
		go notify(addr, remoteIP)
	}
}

func notify(addr, remoteIP string) {
	remote := fmt.Sprintf("%s:%d", addr, notifyport)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", remoteIP)
}

// Servidor de notificaciones
func notifyServer() {
	hostname := fmt.Sprintf("%s:%d", hostaddr, notifyport)
	ln, _ := net.Listen("tcp", hostname)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleNotify(conn)
	}
}

func handleNotify(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	remoteIP, _ := br.ReadString('\n')
	remoteIP = strings.TrimSpace(remoteIP)
	addrs = append(addrs, remoteIP)
}

// Solicitar registro
func requestRegister(remoteIP string) {
	remoteDir := fmt.Sprintf("%s:8000", remoteIP) 	//ip:puerto
	conn, _ := net.Dial("tcp", remoteDir)
	defer conn.Close()
	// Enviar al host remoto mi dirección
	fmt.Fprintf(conn, hostaddr)

	// El host remoto me envia su bitacora de direcciones
	br := bufio.NewReader(conn)
	msg, _ := br.ReadString('\n')

	// Me copio su bitacora y agrego el host remoto
	// Porque el host remoto no se registra en su propia bitacora
	var respAddrs []string
	json.Unmarshal([]byte(msg), &respAddrs)
	
	addrs = append(respAddrs, remoteIP)

	fmt.Println("Pertenesco a la red, mis direcciones son: ", addrs)
}

// Funcion principal
func main() {
	// Nuevo nodo en la red
	hostaddr = getMyIP()
	go registerServer()	// levantar mi servidor de registro
	fmt.Printf("Hola soy %s\n", hostaddr)

	// lectura de puerto remoto
	br := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la ip direccion remota [IP]: ")
	remoteIP, _ := br.ReadString('\n')
	remoteIP = strings.TrimSpace(remoteIP)

	if remoteIP != "" {
		// solicitar registro la solicitud de enrolamiento en la red
		requestRegister(remoteIP)
	}
	notifyServer()	// levantar mi servidor de notificaciones
}
