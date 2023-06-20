package main

import (
	"fmt"
	"sync"
	"time"
)

var memoryAccess sync.Mutex

func P(){
	memoryAccess.Lock()
	fmt.Println("SC P1")
	fmt.Println("SC P2")
	memoryAccess.Unlock()
}

func Q(){
	memoryAccess.Lock()
	fmt.Println("SC Q1")
	memoryAccess.Unlock()

	memoryAccess.Lock()
	fmt.Println("SC Q2")
	memoryAccess.Unlock()
}

func main(){
	go P()
	go Q()
	time.Sleep(100 * time.Millisecond)
}