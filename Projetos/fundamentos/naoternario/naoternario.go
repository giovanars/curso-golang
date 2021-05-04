package main

import (
	"fmt"
)

func obterResultado(nota float64) string{

	if nota >=6{
		return "Aprovado";
	}
	return "Reprovado";
}

func main(){

	
	situacao := obterResultado(5)
	fmt.Println(situacao)
}