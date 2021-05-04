package main

import "fmt"

func media(numeros ...float64) float64 { // o '...' siginificam que essa funcção pode receber quantos parametros quiser

	total := 0.0
	for _, num := range numeros{
		total += num
	}
	return total / float64(len(numeros))
}

func main()  {
	
	fmt.Printf("Media: %.2f", media(7.7, 7.3, 5.3, 8.9, 10))
}