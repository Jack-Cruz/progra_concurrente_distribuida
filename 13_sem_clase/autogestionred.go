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
		if strings.HasPrefix(iface.Name, "Wi-Fi") {
			addrs, _ := iface.Addrs()
			for _, addr := range addrs {
				switch t := addr.(type) {
				case *net.IPNet:
					if t.IP.To4() != nil {
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
	// Nuevo nodo en la red
	miIP := descubrirIP2()
	fmt.Printf("Hola soy %s\n", miIP)

	// lectura de puerto remoto
	br := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la ip direccion remota [IP]")
	remoteIP, _ := br.ReadString('\n')
	remoteIP = strings.TrimSpace(remoteIP)

	if remoteIP != "" {
		// solicitar registro la solicitud de enrolamiento en la red
		requestRegister(remoteIP, miIP)
	}
}

func requestRegister(remoteIP, miIP string) {
	remoteDir := fmt.Sprintf("%s:8000", remoteIP) 	//ip:puerto
	conn, _ := net.Dial("tcp", remoteDir)
	defer conn.Close()
	// Enviar al host remoto la dirección del solicitante
	fmt.Fprintf(conn, miIP)

	// El host remoto va a enviar su bitacora
	br := bufio.NewReader(conn)
	bitacoraRedDest, _ := br.ReadString('\n')

	var bitacoraAux []string
	json.Unmarshal([]byte(bitacoraRedDest), &bitacoraAux)
	// Agregar el nodo a su propia bitacora
	bitacoraAux = append(bitacoraAux, remoteIP)

	fmt.Println(bitacoraAux)
}

func manejadorServicioNuevoHost(conn net.Conn) {
	defer conn.Close()

	// recibir la llamada de nuevo host, ññegar la IP del host
	br := bufio.NewReader(conn)
	remoteIP, _ := br.ReadString('\n')
	remoteIP = strings.TrimSpace(remoteIP)

	// devolver al nuevo host de bitácora
	// serializar
	byteBitacora, _ := json.Marshal(bitacoraRed)
	fmt.Fprintf(conn, "%s\n", string(byteBitacora))
	
	// Notificar al resto de miembro de la red que llegó un nuevo integrante
	for _, dir := bitacoraRed {
		NotificarNuevoHost(dir, remoteIP)
	}
	// Actualizar su bitacora
	bitacoraRed = append(bitacoraRed, remoteIP)
	print("Bitacor actualizada")
}

func NotificarNuevoHost(direccionDestino, remoteIP string) {
	hostDestino := fmt.Sprintf("%s:8001", direccionDestino)
	net.Dial("tcp", hostDestino)
	defer conn.Close()
	fmt.Fprintln(conn, remoteIP)
}