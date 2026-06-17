package main

// * Punteros
// & para obtener la dirección de memoria

var a = 1

// a no es igual a 1, sino a la dirección de memoria donde se almacena el valor 1

func incrementar (numero int) {
	numero++
} // no incrementa el valor de a, sino que incrementa una copia del valor de a

func incrementarConPuntero (numero *int) {
	*numero++ // desreferenciamos el puntero para acceder al valor almacenado en esa dirección de memoria y lo incrementamos
}

func main() {
	valor := a
	incrementar(valor)
	println("Valor después de incrementar sin puntero:", valor) // sigue siendo 1
	incrementarConPuntero(&a)
	println("Valor después de incrementar con puntero:", a) // ahora es 2
}