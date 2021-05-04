package main

import (
	"fmt"
)

func main()  {
	funcsESalarios := map[string]float64{
		"José": 11325.45,
		"Gabiela": 15456.70,
		"Pedro": 1200,
	}

	funcsESalarios["Rafael"] = 1350
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios{
		fmt.Println(nome, salario)
	}
}