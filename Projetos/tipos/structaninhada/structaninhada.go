package main

import "fmt"

type item struct {
	produtoID int
	qtd int
	preco float64
}

type pedido struct {
	userID int
	items []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0

	for _, item := range p.items {
		total += item.preco * float64(item.qtd)
	}

	return total;
}

func main()  {
	pedido := pedido{
		userID: 1,
		items: []item {
			item{1, 2, 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	fmt.Printf("Valor total do prdido é %.2f", pedido.valorTotal())
}