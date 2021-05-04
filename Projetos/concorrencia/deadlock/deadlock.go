package main

import "fmt"
import "time"

func rotina(c chan int){
	time.Sleep(time.Second)
	c <- 1 //operação bloaqueante
	fmt.Println("Só  depois que foi lido")
}

func main()  {
	c := make(chan int) //canal sem buffer
	go rotina(c)
	fmt.Println(<-c) //operação bloaqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) //deadlock
	fmt.Println("Fim")
}