package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Ingrese un número:")
	bufferIn := bufio.NewReader(os.Stdin)
	ingreso, err := bufferIn.ReadString('\n')

	// Manejo de excepciones (error)
	if err != nil {
		fmt.Println("Mensaje de error 1: ", err.Error())
		os.Exit(1)
	}

	ingreso = strings.Trim(ingreso, "\r\n")
	num, er := strconv.Atoi(ingreso)

	// Manejo de excepciones (error)
	if er != nil {
		fmt.Println("Mensaje de error 2: ", er.Error())
		os.Exit(1)
	} 
	
	fmt.Printf("El doble de %v es %v", num, 2*num)
}