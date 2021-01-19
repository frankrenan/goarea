package area

import "math"

//Pi é um constante
const Pi = 3.15

//Circ é uma circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula o retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

//Não vísivel
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
