package main

import "fmt"

func inc1(n int){
	n++
}

//ponteiro é representa por *
func inc2(n *int){
	//* é usado para desrefernciar, ou seja, ter acesso a valor no qual o ponteiro aponta
	*n++;
}

func main()  {
	n := 1

	inc1(n)
	fmt.Println(n)

	inc2(&n)//pegando o endereço na memoria de n e passando para funcção e por isso o valor de n vai mudar 
	fmt.Println(n)
}