package main

import "fmt"
import "strconv"

func main()  {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	//cuidado
	fmt.Println("Teste " + string(123))


	//int to string
	fmt.Println("Teste " + strconv.Itoa(123))

	//string to int -- Função que retorna dois valores, 1º valor convertido e o 2º é um erro caso não de pra converter
	//_ é como ignorar a variavel porque se eu declara algo eu vou ter que usar
	num, _ := strconv.Atoi("123");

	fmt.Println(num - 122);

	//string to bool -- funciona no mesmo esquema da de cima
	b, _ := strconv.ParseBool("true");

	if b{
		fmt.Println("Verdadeiro");
	}
}