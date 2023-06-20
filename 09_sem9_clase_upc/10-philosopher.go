package main

import "fmt"

func philosopher(name string, rightFork, leftFork chan bool) {
	for {
		fmt.Printf("%s esta pensando\n", name)
		<-rightFork
		<-leftFork
		fmt.Printf("%s esta comiendo\n", name)
		rightFork <- true
		leftFork <- true
	}
}


func fork(ithFork chan bool) {
	for {
		ithFork <- true
		<- ithFork
	}
}

func main() {
	n := 5
	forks := make([]chan bool, n)
	names := []string{"Descartes", "Niezche", "Socrates", "Aristoteles"}

	for i := range forks {
		forks[i] = make(chan bool, 1)
	}
	for i := 0; i < n-1; i++ {
		go philosopher(names[i], forks[i], forks[i+1])
		go fork(forks[i])
	}
	go fork(forks[n-1])
	philosopher("Susy", forks[n-1], forks[0])

}