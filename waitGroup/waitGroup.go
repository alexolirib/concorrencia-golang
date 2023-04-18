package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	// informo quantos goroutine deve ser esperado a execucao
	waitGroup.Add(2)

	go func ()  {
		write("Olá mundo")
		waitGroup.Done()
	}()
	go func ()  {
		write("Programando em Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func write(txt string) {
	for i := 0; i < 5; i ++{
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}