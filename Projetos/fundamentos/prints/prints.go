package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")

	fmt.Println(" Nova")
	fmt.Println(" linha.")

	x := 3.141516

	//fmt.Println("O valor de x é " + x) //Isso não funciona

	xs := fmt.Sprint(x)

	fmt.Println("O valor de x é " + xs) //Concatenação de strigs

	fmt.Println("O valor de x é", x) //Tbm funciona

	fmt.Printf("O valor de x é %.2f", x) //Aqui é possivel formatar

	a, b, c, d := 1, 1.999, false, "opa"

	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)

}
