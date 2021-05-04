package main

import "fmt"

func main(){
	funcsPorLetra := map[string]map[string]float64{
		"G":{
			"Gabi": 156464.65,
			"Guga": 846546,
		},
		"J":{
			"José": 15455,
		},
		"P":{
			"Pedro": 1236,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra{
		fmt.Println(letra)
		for nome, salario := range funcs{
			fmt.Println(nome, salario)
		}
	}
}