package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo"); //go routines - não espera termina para continuar execução
	escrever("programando em go");
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
