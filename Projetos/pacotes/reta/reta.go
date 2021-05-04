package main

import (
	"math"
)

//Inicializando um atributo com a letre MAIUSCULA vc cria algo publico (visivel fora e dentro do pacote)
//Iniando com a letra MINUSCULA é privado (dentro do pacote e não somente no arquivo)
//No go não existi visibilidade PRIVADA dentro de arquivo e sim dentro do pacote

//Ponto representa uma cordenado plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.x * a.x)
	cy = math.Abs(b.y * a.y)
	return
}

//Distancia é resposavel por calcula a distancia linear entre dois pontp
func Distancia(a, b Ponto) float64{
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx,2) + math.Pow(cy, 2))
}
