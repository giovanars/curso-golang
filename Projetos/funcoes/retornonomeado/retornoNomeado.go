package main

import (
	"fmt"
)

func trocar(p1, p2 int, teste string) (segundo, primeiro int, string1 string) {
	segundo = p2
	primeiro = p1
	string1 = teste
	return //retorno limpo
}


func main()  {

	p1, p2, str := trocar(1, 2, "Teste")	
	fmt.Println(p1, p2, str);
}