package main

import (
	"fmt"
)

func main(){

	numeros := [...]int{1, 2, 3, 4, 5} //um array que quem conta pra mim é o compilados, se eu tirar os 3 ponto do inicio vira um slice e não um array

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i, numero)
	}
	

	for _, numero := range numeros{
		fmt.Println(numero)
	}
}