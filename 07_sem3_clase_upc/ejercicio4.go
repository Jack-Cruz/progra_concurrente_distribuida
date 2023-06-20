package main

import (
	"fmt"
	"sync"
)

var memoryAccess sync.Mutex

var value int

func main() {
	// funciones anónimas
	go func() {
		memoryAccess.Lock()
		value++					// Sección crítica
		memoryAccess.Unlock()
	}()
	
	
	for i := 0; i < 10; i++ {
		if value == 0 {
			fmt.Printf("El valor es %v\n", value)
		} else {
			fmt.Printf("El valor es %v\n", value)
		}
	}
}

