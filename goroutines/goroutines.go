package main

import (
	"fmt"
	"time"
)

// concorrencia != paralelismo
func main() {
	// deixa os processos em concorrencia 
	go write("Olá mundo")
	write("Programando em Go")
}

func write(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second*2)
	}
}