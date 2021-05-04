package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int){
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (interação %d)\n", pessoa, texto, i+1)
	}
}


func main(){
	//fale("Maria", "Porque vc não fala comigo?", 3)
	//fale("João", "Só posso falar depois de vc", 1)

	// go fale("Maria", "Ei...", 500) //só continua executando se a thread principal não terminar
	// go fale("João", "Opa...", 500)
	// fmt.Println("Fim!")

	go fale("Maria", "Entendi", 10)
	fale("João", "Parabéns", 5)
}