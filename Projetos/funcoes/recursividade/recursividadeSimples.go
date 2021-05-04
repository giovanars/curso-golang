package main

import "fmt"

func fatorial(n uint) uint {//recebendo inteiro possitivo e retornando um inteiro possitvo
	
	switch {
	case n == 0:
		return 1
	default:
		fatorialAnterior := fatorial(n - 1)
		return n * fatorialAnterior
	}
}

func main()  {
	resultado := fatorial(5)
	fmt.Println(resultado)
}