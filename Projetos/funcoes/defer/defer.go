package main

import "fmt"

func obterValorAprovado(numero int) int{
	defer fmt.Println("Fim")//executa no ultimo momento antes de sair da função, nesse caso antes do return
	if numero >= 5000 {
		fmt.Println("Valor alto")
		return 5000
	}else{
		fmt.Println("Valor baixo")
		return numero
	}
}

func main()  {

	fmt.Println(obterValorAprovado(6000))
	fmt.Println(obterValorAprovado(4000))
}