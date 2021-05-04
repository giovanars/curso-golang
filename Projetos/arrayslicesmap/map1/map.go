package main

import (
	"fmt"
)

func main()  {
	//var aprovados map[int]string
	//maps devem ser anicializados

	aprovados := make(map[int]string)

	aprovados[12345678978] = "Maria"
	aprovados[16465465544] = "Pedro"
	aprovados[98765432112] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[98765432112])
	delete(aprovados,98765432112)
	fmt.Println(aprovados[98765432112])


}