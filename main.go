package goarea

import "math"

// quando é publico tem que ter comentario associado - pi é valor de variavel
const Pi = 3.14

// circ calcula area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// rect calcula area retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// nao visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
