package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func player(id int, wg *sync.WaitGroup, in, out chan int) {
    defer wg.Done()
    for {
        // Esperar a recibir el aro del jugador anterior
        //<-in

        // Realizar el valor aleatorio
        rand.Seed(time.Now().UnixNano())
        randomValue := rand.Intn(100)

        // Imprimir el valor generado
        fmt.Printf("Jugador %d: Valor aleatorio generado: %d\n", id, randomValue)

        // Comparar con el jugador siguiente
        nextValue := <-in
        if randomValue > nextValue {
            // Pasar el aro al siguiente jugador
            fmt.Printf("Jugador %d: Valor mayor, pasando el aro\n", id)
            out <- randomValue
        } else {
            // Volver a empezar si el valor es menor o igual
            fmt.Printf("Jugador %d: Valor menor o igual, reiniciando\n", id)
            out <- nextValue
            continue
        }
    }
}

func main() {
    // Número de jugadores y aros disponibles
    numPlayers := 4

    // Crear canales para comunicarse entre jugadores
    players := make([]chan int, numPlayers)
    for i := range players {
        players[i] = make(chan int)
    }

    // Crear WaitGroup para esperar a que todos los jugadores terminen
    var wg sync.WaitGroup
    wg.Add(numPlayers)

    // Iniciar el juego
    for i := 0; i < numPlayers; i++ {
        go player(i+1, &wg, players[i], players[(i+1)%numPlayers])
    }

    // Generar el primer valor aleatorio
    rand.Seed(time.Now().UnixNano())
    firstValue := rand.Intn(100)
    fmt.Printf("Valor aleatorio generado por el primer jugador: %d\n", firstValue)

    // Comenzar el juego pasando el primer valor aleatorio al primer jugador
    players[0] <- firstValue

    // Esperar a que todos los jugadores terminen
    wg.Wait()
}