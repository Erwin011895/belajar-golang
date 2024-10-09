package main

import (
	"fmt"
)

type BisaMemotong interface {
	Cut()
	SliceNTimes(n int)
}

type scissor struct {
	brand string
}

func (s scissor) Cut() {
	fmt.Println(s.brand, "is cutting..")
}

func (s scissor) SliceNTimes(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(s.brand, "is cutting..")
	}
}

type knife struct {
	brand string
}

func (k knife) Cut() {
	fmt.Println(k.brand, "is used")
}

func (k knife) SliceNTimes(n int) {
	fmt.Println(k.brand, "is used", n, "times")
}

type hammer struct {
	brand string
}

// see another struct definition on model/programmer & model/company

func main() {
	// interface basic
	guntingJoyko := scissor{
		brand: "Joyko",
	}

	var objekBisaPrint BisaMemotong
	objekBisaPrint = guntingJoyko
	objekBisaPrint.Cut()
	objekBisaPrint.SliceNTimes(2)

	fmt.Println("==========")

	pisauStein := knife{brand: "STEIN"}

	CustomCut(guntingJoyko)
	CustomCut(pisauStein)

	// paluEnamel := hammer{brand: "Enamel"}
	// CustomCut(paluEnamel) // error
}

func CustomCut(obj BisaMemotong) {
	fmt.Println("ini custom cut")
	obj.Cut()
	obj.Cut()
	fmt.Println("custom cut selesai")
}
