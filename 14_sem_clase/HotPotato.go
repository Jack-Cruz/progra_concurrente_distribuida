package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// Bitacora de direcciones de miembros de la red
// menos el mismo
var bitacoraRed[] string

// direccion de red del nodo
var dir_nodo string

// Puertos de servicio del nodo
const (
	puerto_registro = 8000	// autogestion
	puerto_notificacion = 8001 	// autogestion
	puerto_proceso = 8002	// servicio (enviar, escuchar)
	puerto_solicitud = 8003	// awrawala
)

// Estructuras de mensajes
type Info struct {
	Tipo string
	NodeNum int
	NodeDir string
}

type MyInfo struct {
	contMsg int
	primero bool
	nextNum int
	nextDir string
}

var puedeIniciar chan bool
var chMyInfo chan MyInfo
var ticket int

func main() {
	// Indentificarse
	dir_nodo = obtenerIP()

	// Rol de servidor (Escucha)
	// Para registrar nuevos nodos
	go registrarServidor()

	// Servicio del nodo (Esucha)
	// Hot potato
	go registrarServicioHP()

	// Rol de cliente (Solicitar unirse a la red - Autogestionarse)
	br := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese la ip del nodo remoto a conectarse: ")
	remotehost, _ := br.ReadString('\n')
	remotehost = strings.TrimSpace(remotehost)
	
	// Para el primer nodo que recien inicia en la red
	if remotehost != "" {
		registrarCliente(remotehost)
	}
	// rol del servidor
	escucharNotificaciones()

	// Generar el ticket
	rand.Seed(time.Now().UTC().UnixNano())
	ticket = rand.Intn(1000000)
	fmt.Println("Ticket: ", ticket)

	// Crear los canales
	puedeIniciar = make(chan bool)
	chMyInfo = make(chan MyInfo)

	// enviar la solicitud inicial
	go func() {
		chMyInfo <- MyInfo{0, true, int(1e7), ""}
	}()

	// Esperar el inicio de la solicitud: cliente envia
	go func() {
		fmt.Print("Presione enter para iniciar la solicitud...")
		br := bufio.NewReader(os.Stdin)
		br.ReadString('\n')

		// Crear el mensaje de solicitud
		info := Info{"SENDNUM", ticket, dir_nodo}

		// Notificar el mensaje de la solicitud
		for _, dir := range bitacoraRed {
			send(dir, info)
		}
	}()
	
	// Modo escucha para recibir las solicitudes
	AtenderSolicitudes()
}



// Metodos de autogestion

// Obtener mi ip
func obtenerIP() string {
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
	return "127.0.0.1"
}

// Servidor de registro (modo escucha)
func registrarServidor() {
	hostname := fmt.Sprintf("%s:%d", dir_nodo, puerto_registro)
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
	fmt.Println("Recibi un nuevo registro")
	remoteIP, _ := br.ReadString('\n')
	remoteIP = strings.TrimSpace(remoteIP)
	fmt.Printf("Nuevo host en la red: %s\n", remoteIP)
	
	// Devolver al nuevo host, la bitacora que guardo
	// Enviar mi bitacora de addrs al nuevo host
	addrbytes, _ := json.Marshal(bitacoraRed)
	fmt.Fprintf(conn, "%s\n", string(addrbytes))
	
	// Notificar a todos que llego un nuevo host
	notificarTodos(remoteIP)

	// Agregar el nuevo host a mi bitacora
	bitacoraRed = append(bitacoraRed, remoteIP)
}

func notificarTodos(remoteIP string) {
	for _, dir := range bitacoraRed {
		go notify(dir, remoteIP)
	}
}

func notify(dir, remoteIP string) {
	remote := fmt.Sprintf("%s:%d", dir, puerto_notificacion)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	fmt.Fprintf(conn, "%s\n", remoteIP)
}

// Solicitar registro
func registrarCliente(remoteIP string) {
	// se usa cuando ya existe una red y se quiere unir el nodo
	remoteDir := fmt.Sprintf("%s:%d", remoteIP, puerto_registro) 	//ip:puerto
	fmt.Println("Solicito conexion a", remoteDir)
	conn, _ := net.Dial("tcp", remoteDir)
	defer conn.Close()
	// Enviar al host remoto (DATA => mi dirección)
	fmt.Fprintf(conn, "%s\n", dir_nodo)
	fmt.Println("Envie mi host")

	// El host remoto me envia su bitacora de direcciones
	br := bufio.NewReader(conn)
	msg, _ := br.ReadString('\n')
	fmt.Println("Recibi sus direcciones")

	// Me copio su bitacora y agrego el host remoto
	// Porque el host remoto no se registra en su propia bitacora
	var respAddrs []string
	json.Unmarshal([]byte(msg), &respAddrs)
	
	bitacoraRed = append(respAddrs, remoteIP)

	fmt.Println("Pertenesco a la red, mis direcciones son: ", addrs)
}

// Servidor de notificaciones
func escucharNotificaciones() {
	hostname := fmt.Sprintf("%s:%d", dir_nodo, puerto_notificacion)
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

	// cada nodo va a registrar la ip en su bitacora
	bitacoraRed = append(bitacoraRed, remoteIP)
	fmt.Println("Bitacora cliente notificado", bitacoraRed)
}




// Servidor de Hot Potato

func registrarServicioHP() {
	hostname := fmt.Sprintf("%s:%d", dir_nodo, puerto_proceso)
	ln, _ := net.Listen("tcp", hostname)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go handleRegisterHP(conn)
	}
}

func handleRegisterHP(conn net.Conn) {
	defer conn.Close()

	// recibir la llamada del nuevo host, llego la papa caliente
	br := bufio.NewReader(conn)
	strNum, _ := br.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(strNum))
	fmt.Println("Recibimos el %d\n", num)

	if num == 0 {
		fmt.Println("Finalizo el proceso, game over!!!")
	} else {
		enviarProximoNodo(num-1)
	}
}

func enviarProximoNodo(num int) {
	// Seleccionar de forma aleatoria el proximo nodo
	idx := rand.Intn(len(bitacoraRed))
	// mensaje
	fmt.Printf("Enviando el %d a %s\n", num, bitacoraRed[idx])
	// formatear la direccion del proximo nodo
	remoteHost := fmt.Sprintf("%s:%d", bitacoraRed[idx], puerto_proceso)
	conn, _ := net.Dial("tcp", remoteHost)
	defer conn.Close()
	fmt.Fprintf(conn, "%d\n", num)
}



// Metodos de tickets (agrawala)

func send(remoteAddr string, info Info) {
	remote := fmt.Sprintf("%s:%d", remoteAddr, puerto_solicitud)
	conn, _ := net.Dial("tcp", remote)
	defer conn.Close()
	bytesMsg, _ := json.Marshal(info)
	fmt.Fprintln(conn, string(bytesMsg))
}

func AtenderSolicitudes() {
	host := fmt.Sprintf("%s:%d", dir_nodo, puerto_solicitud)
	ln, _ := net.Listen("tcp", host)
	defer ln.Close()
	for {
		conn, _ := ln.Accept()
		go HandleSolicitud(conn)
	}
}

func HandleSolicitud(conn net.Conn) {
	defer conn.Close()
	br := bufio.NewReader(conn)
	msg, _ := br.ReadString('\n')
	var info Info
	json.Unmarshal([]byte(msg), &info)
	fmt.Println("Recibi siguiente info ", info)
	
	switch info.Tipo {
	case "SOLICITUD":
		myInfo := <- chMyInfo
		myInfo.contMsg++

		if info.NodeNum < ticket {
			myInfo.primero = false
		} else if info.NodeNum < myInfo.nextNum {
			myInfo.nextNum = info.NodeNum
			myInfo.nextAddr = info.nextDir
		}


		go func() {
			chMyInfo <- myInfo
		}()
		if myInfo.contMsg == len(addrs) {
			if myInfo.primero {
				fmt.Println("Soy el primer!! :)")
				criticalSection()
			} else {
				puedeIniciar <- true
			}
		}
	case "START":
		<- readyToStart
		criticalSection()
	}
}

func criticalSection() {
	fmt.Println("Critical section")
}