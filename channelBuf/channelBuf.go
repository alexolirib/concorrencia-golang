package main

import "fmt"

func main() { 
	// buffer - 2 só vai bloquear quando chegar na capacidade máxima asism é possível utlizar 
	// channel no mesmo método
	// se passar do buffer da deadlock
	channel := make(chan string, 2)
	channel <- "Olá mundo"
	channel <- "Olá mundo2"
	// channel <- "Olá mundo3"

	message := <-channel
	message2 := <-channel
	// message3 := <-channel
	fmt.Println(message)
	fmt.Println(message2)
	// fmt.Println(message3)
}