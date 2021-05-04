package main

import (
	"fmt"
)

type imprimivel interface{
	toString() string
}

type pessoa struct{
	nome string
	sobrenome string
}

type produto struct{
	nome string
	preco float64
}

//interfaces são implementas implicitamente

func (p pessoa) toString() string{
	return p.nome + " " + p.sobrenome
}

func (p produto) toString() string{
	return fmt.Sprintf("%s - R$ %.2f", p.nome, p.preco)
}


//no go não existe algum comando que faça uma objeto implementas o outro, pra ele reconhece um interface basta ter os metodos da struct que respeitem o contrato
func imprimir(x imprimivel){
	fmt.Println(x.toString())
}

func main()  {
	var coisa imprimivel = pessoa{"Roberto", "Silva"}
	fmt.Println(coisa.toString())
	imprimir(coisa)

	coisa = produto{"Caneta", 10}//jogando o produto dentro de coisa podemos ver aqui o polimorfismo, onde uma coisa pode ser varias. Essa coisa foi primeiro uma pessoa e depois um produto
	fmt.Println(coisa.toString())
	imprimir(coisa)
}