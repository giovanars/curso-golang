package main

import "fmt"

func init(){//função que é executada antes da main, mesmo não chamando ela na main ele vai rodar
	fmt.Println("Inicialização...")
}

func main()  {
	fmt.Println("Main...")
}