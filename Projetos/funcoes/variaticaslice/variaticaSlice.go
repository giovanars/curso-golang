package main

import "fmt"

func imprimirAprovados(aprovados ...string){

	fmt.Println("Lista de Aprovados")

	for i, aprovado := range aprovados{
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}

}

func main()  {
	
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilerme"}
	imprimirAprovados(aprovados...)//usando os 3 pontinhos o compilados quebra o slice como se fosse parametros
}