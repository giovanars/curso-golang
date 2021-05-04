package main 

import "fmt"
import "time"

func main(){

	i :=1
	for i <= 10{
		fmt.Printf("%d ", i)
		i++
	}

	for i:= 0; i <= 20; i++ {
		fmt.Printf("%d ", i)
	}

	for i := 0; i <=10; i++ {
		if i%2 == 0{
			fmt.Println("Par")
		} else {
			fmt.Println("Ímpar")
		}
	
	}

	for {
		//Loop infinito

		fmt.Println("Para sempre ...")
		time.Sleep(time.Second)
	}
}