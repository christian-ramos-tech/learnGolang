package main //package name
// package	otroPaquete 

import "fmt" // librerías necesarias para importar el paquete

func main() {
	// Variables
	var mensaje string = "Hola mundo"
	easyMessage := "Hola mundo usando :="
	fmt.Println(mensaje)
	fmt.Println(easyMessage)

	// Puedes comentar usando "//"

	// float numbers
	a := 10.
	var b float64 = 3
	fmt.Println(a / b)


	//integer numbers
	var c int = 10
	d := 3
	fmt.Println(c / d)

	// boolean
	var x bool = true
	y := false
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(!x)

	//Lógica privada
	privada()

	//Lógica Publica
	Publica()

	//Función "defer"
	printHelloWorld()
}

// func types. Si empieza con minuscula es privada y solo puede llamarse en este caso desde main.go
func privada() {
	fmt.Println("Ejecutar lógica que no necesita ser exportada (pertenece solo a este paquete)")
}


// si empieza con mayúscula hace que la función sea pública
func Publica() {
	fmt.Println("Lógica que quiero exportar a otros paquetes")
}

// defer hace que esa línea de código se ejecute al final
func printHelloWorld() {
	defer fmt.Println("World!")
	fmt.Println("Hello")
	fmt.Println("defer ejecuta un fragmento de código hasta que la función haya terminado")
}
