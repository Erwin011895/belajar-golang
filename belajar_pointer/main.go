package main

import "fmt"

func main() {
	angka := 10
	var pointerKeAngka *int = &angka

	// explain pointer
	fmt.Println("angka", angka)
	fmt.Println("&angka", &angka)
	fmt.Println("pointerKeAngka", pointerKeAngka)
	fmt.Println("*pointerKeAngka", *pointerKeAngka)
	fmt.Println("==========")

	// modify common variable
	angka = 99
	fmt.Println("angka", angka)
	fmt.Println("*pointerKeAngka", *pointerKeAngka) // value di *pointer ikut keganti
	fmt.Println("==========")

	// modify *pointer value
	*pointerKeAngka = 50
	fmt.Println("angka", angka) // value di variable ikut keganti
	fmt.Println("*pointerKeAngka", *pointerKeAngka)
	fmt.Println("==========")

	fmt.Println("before", angka)

	nonPointerFunc(angka)
	fmt.Println("after nonPointerFunc", angka)

	pointerFunc(&angka)
	fmt.Println("after pointerFunc", angka)

	angka = 0 // set to 0
	fmt.Println("pointerFunc ke-2", angka)

	pointerFunc(pointerKeAngka)
	fmt.Println("after pointerFunc ke-2", angka)
}

func nonPointerFunc(x int) {
	x = x + 10
}

func pointerFunc(x *int) {
	*x = 777
}
