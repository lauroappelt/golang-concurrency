package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // 2 goroutines nesse grupo de espera

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("programando em go")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // espera as go routines acabarem
}

func escrever(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
