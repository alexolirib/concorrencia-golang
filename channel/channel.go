package main

import (
	"fmt"
	"time"
)

func main() {
	// trafegar dados do tipo string
	channel := make(chan string)
	go write("olá mundo", channel)

	fmt.Println("Depois da funcao write")
	// for {
	// 	//espera que seja atribuido 
	// 	message, open := <-channel
	// 	//valida se o canal ta aberto
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }
	// forma mais resumida - finaliza quando o canal fecha
	for message := range channel {
		fmt.Println(message)
	}
	fmt.Println("Fim do programa")
}

func write(txt string, channel chan string) {
	time.Sleep(time.Second)
	for i := 0; i < 5; i ++{
		channel <- txt
		time.Sleep(time.Second)
	}
	//fecha o canal (importante para nao ocorrer deadlock)
	close(channel)
}