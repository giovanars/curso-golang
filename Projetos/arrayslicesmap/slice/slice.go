package main

import "fmt"

func main()  {
	//Slice não é um array, ele define um pedaço de um array
	
	a1 := [3]int{1,2,3} //array
	s1 := []int{1,2,3}//slice
	fmt.Println(a1, s1)

	a2 := [5]int{1,2,3,4,5}
	s2 := a2[1:3]//isso não gera um array diferente, ele tem um ponteiro para o elemento para o qual ela aponta
	fmt.Println(a2, s2)

	s3 := a2[:2]
	fmt.Println(a2, s3)

	s4 := s2[:1]
	fmt.Println(s2, s4)

	//todos os slices aponta pra o mesmo array criado na linhado 12

}