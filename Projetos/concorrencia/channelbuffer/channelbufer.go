package main

import (
	"fmt"
	"time"
)


func rotina(c chan int){
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	fmt.Println("Executou")
	c <- 6
}

func main()  {
	ch := make(chan int, 3)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	// não é obrigatorio ler todos os dados de um canal eu só não posso ler a mais senão vai dar deadlock
}