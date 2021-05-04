package main

import (
	"fmt"
)

func main(){
	i := 1

	//GO não tem aritimetica com ponteiro

	var p *int = nil

	p = &i //pegando a referencia da variavel
	*p++   //desreferenciando (pegando valor)
	i++

	fmt.Println(p, *p, i, &i)
}