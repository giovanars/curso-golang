package main

//Channel (canal) é a forma de comunicação entre gorotines
//channel é um tipo

import (
	"fmt"
	"time"
)

func doisTresQuatroVezes(base int, c chan int)  {
	time.Sleep(time.Second)
	c <- 2 * base

	time.Sleep(time.Second)
	c <- 3 * base

	time.Sleep(3 * time.Second)
	c <- 4 * base
}

func main()  {
	c := make(chan int)
	go doisTresQuatroVezes(2, c) //go routine causa o assincronisno e o channel causa o sincronismos, ele é que espera 

	a, b := <-c, <-c //recebendo dados do canal
	fmt.Println(a, b)
	fmt.Println(<-c)

}