package main

import (
	"net/http"
	"log"
)

func main(){
	http.HandleFunc("/usuarios", UsuarioHandler)
	log.Println("Funcionando...")
	log.Fatal(http.ListenAndServe(":3000", nil))

}