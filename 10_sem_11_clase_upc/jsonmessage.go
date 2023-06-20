package main

import (
	"encoding/json"
	"fmt"
)

type Alumno struct {
	Codigo string `json:"cod"`
	Nombre string `json: "nom"`
	Promedio float32 `json: "prom"`
}

func main() {
	alumnos := []Alumno{
		{"u201911123", "Juan Garcia", 12.49},
		{"u201917514", "Maria Gonazales", 12.49},
		{"u201912400", "Juan Garcia", 12.49},
	}

	// Serializar (estructura en memoria -> arreglo de bytes -> string)
	jsonBytes, _ := json.MarshalIndent(alumnos, "", "  ")
	jsonStr := string(jsonBytes)
	fmt.Println("JSON", jsonStr)

	// Deserializar (string -> arreglo de bytes -> estructura en memoria)
	var alumnos2 []Alumno
	json.Unmarshal(jsonBytes, &alumnos2)
	fmt.Println("Struct", alumnos2)
}