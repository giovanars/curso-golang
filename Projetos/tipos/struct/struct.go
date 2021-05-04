package main

import "fmt"

type produto struct {
	nome string
	preco float64
	desconto float64	
}


//Metodo: função com receiver (repector)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main()  {
	var p1 produto
	p1 = produto{
		nome: "Caneta",
		preco: 1.55,
		desconto: 0.05,
	}
	fmt.Println(p1, p1.precoComDesconto())

	p2 := produto{"Notebook", 1998.09, 0.10}
	fmt.Println(p2.precoComDesconto())
}