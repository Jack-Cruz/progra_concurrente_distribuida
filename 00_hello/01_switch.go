package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
	fmt.Println("Time to relax")
	fmt.Println(randomQuote())
	fmt.Println(randomSong())
}

func randomQuote() string {
	quotes := []string {
		"Cuando hacemos un proyecto, debemos documentarlo!",
		"Debemos ser prácticos en ciertas circunstancias",
		"Antes de graduarme quiero completar un libro de programación -> Cormen",
		"Todo en exceso es dañino, lo más bonito es el punto medio",
		"La vida no es ordenada, perfecta, secuencial, sobretodo en los estudios",
		"Impotencia de no mejorar, tener a cargo a Abel, empezar de nuevo",
	}
	return quotes[rand.Intn(len(quotes))]
}

func randomSong() string {
	songs := []string {
		"Trascender de Mauricio Alen",
		"Hasta la Raíz de Natalia Lafourcade",
		"De Ellos Aprendi de David Rees",
		"Abracadabra de Ami Rodriguez",
		"Festival de Los Polinesios, RedOne",
		"Gracias de Los Polinesios",
		"Contigo quiero estar de Skabeches",
		"La Partida de Antrax",
		"Sonrisas de Pascal",
		"Frio frio de Juan Luis Guerra y Romeo Santos",
		"Tu Foto de Ozuna",
		"Me Voy Enamorando de Chino & Nacho",
		"Rude de MAGIC!",
		"See You Again de Wiz Khalifa, Charlie Puth",
		"My Heart Will Go On de Céline Dion, James Horner",
		"Savage Love de Jawsh 685, Jason Derulo",
		"A Thousand Years de Cristina Perri",
		"Speechless de Naomi Scott",
		"Hello de Adele",
		"Perfect de Ed Sheeran",
	}
	return songs[rand.Intn(len(songs))]
}