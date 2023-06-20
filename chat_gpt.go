package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Enumeración para las opciones de Rock Paper Scissors
const (
	Rock     = 0
	Paper    = 1
	Scissors = 2
)

// Estructura de un jugador
type Player struct {
	ID     int
	Team   int
	Tokens int
}

// Función para generar una jugada aleatoria
func generateMove() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3)
}

// Función para jugar Rock Paper Scissors
func playRPS(player1, player2 *Player) int {
	move1 := generateMove()
	move2 := generateMove()

	if move1 == move2 {
		return 0 // Empate
	} else if (move1 == Rock && move2 == Scissors) ||
		(move1 == Paper && move2 == Rock) ||
		(move1 == Scissors && move2 == Paper) {
		return player1.ID // Jugador 1 gana
	} else {
		return player2.ID // Jugador 2 gana
	}
}

// Función para que un jugador realice su turno
func playerTurn(player *Player, hoops chan int, tokens chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		hoop := <-hoops // Esperar a que haya un aro disponible
		fmt.Printf("Player %d from Team %d is jumping through hoop %d\n", player.ID, player.Team, hoop)
		time.Sleep(time.Second) // Simular tiempo de salto

		// Verificar si el jugador llegó a un cono de otro equipo
		if hoop != player.Team {
			player.Tokens++
			fmt.Printf("Player %d from Team %d collected a token!\n", player.ID, player.Team)
			tokens <- player.Team // Enviar ficha al equipo correspondiente
			break
		} else {
			hoops <- hoop // Volver a colocar el aro en el canal
			opponent := <-tokens // Esperar a que haya una ficha disponible para jugar contra otro jugador
			if player.Team != opponent {
				// Jugar Rock Paper Scissors con el oponente
				winner := playRPS(player, &Player{ID: opponent})
				if winner == player.ID {
					fmt.Printf("Player %d from Team %d won Rock Paper Scissors!\n", player.ID, player.Team)
				} else if winner == opponent {
					fmt.Printf("Player %d from Team %d lost Rock Paper Scissors!\n", player.ID, player.Team)
					break
				}
			}
		}
	}

}

func main() {
	numTeams := 6
	teamSize := 3

	hoops := make(chan int, numTeams)   // Canales para los aros disponibles
	tokens := make(chan int, numTeams)  // Canales para las fichas recolectadas por cada equipo
	waitGroup := sync.WaitGroup{}       // WaitGroup para sincronizar los goroutines
	playerCounter := 1                  // Contador para asignar ID a cada jugador

	// Inicializar los canales y aros disponibles
	for i := 1; i <= numTeams; i++ {
		hoops <- i
		tokens <- i
	}

	// Crear y ejecutar los goroutines para los jugadores
	for team := 1; team <= numTeams; team++ {
		for player := 1; player <= teamSize; player++ {
			waitGroup.Add(1)
			go playerTurn(&Player{ID: playerCounter, Team: team}, hoops, tokens, &waitGroup)
			playerCounter++
		}
	}

	waitGroup.Wait() // Esperar a que todos los jugadores terminen

	fmt.Println("Game over!")
}
