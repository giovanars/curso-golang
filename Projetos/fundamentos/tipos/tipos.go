package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	///inteiros
	fmt.Println(1, 2, 100)
	fmt.Println("Tipo literal inteiro é ", reflect.TypeOf(32000))

	//uint8
	var b byte = 255
	fmt.Println("o byte é ", reflect.TypeOf((b)))


	i1 := math.MaxInt64
	fmt.Println("o valor máximo do int é ", i1)
	fmt.Println("o tipo do valor máximo do int é ", reflect.TypeOf((i1)))

	//rune, valor inteiro(int32) da tabela NCODE
	var i2 rune = 'a'
	fmt.Println("o rune é ", reflect.TypeOf((i2)))
	fmt.Println("o valor do rune é ", i2)

	//reais
	var x float32 = 49.99
	fmt.Println("o tipo de x é ", reflect.TypeOf(x))
	fmt.Println("o tipo literal de x é ", reflect.TypeOf(49.99))

	//boolean
	bo := true
	fmt.Println("o tipo de bo é", bo)
	fmt.Println(!bo)

	//string

	s1 := "olá meu nome é Gi"

	fmt.Println(s1 + "!")
	fmt.Println("o tamanho de s1 é ", len(s1))

	s2 := `Olá
	meu 
	nome
	é
	Gi`

	fmt.Println("o tamanho de s2 é ", len(s2))

	//char???Não existe, guarda o valor int32 de referencia da tabela NCODE;
	char := 'a'
	fmt.Println("o tipo de char é ", reflect.TypeOf(char))
}
