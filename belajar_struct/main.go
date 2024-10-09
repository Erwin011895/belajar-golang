package main

import (
	"fmt"

	"github.com/Erwin011895/belajar-golang/belajar_struct/model"
)

// declare struct type
type person struct {
	name string
	age  int
}

// see another struct definition on model/programmer & model/company

func main() {
	orang := person{}
	fmt.Println("orang:", orang)

	// set value
	orang.name = "Erwin"
	orang.age = 25

	// get value
	fmt.Println("name:", orang.name)
	fmt.Println("age:", orang.age)

	// literal init struct
	orangKedua := person{
		name: "Einstein",
		age:  76,
	}

	// get value for operations
	orang.age++
	if orang.age < orangKedua.age {
		fmt.Println(orang.name, "lebih muda")
	} else {
		fmt.Println(orangKedua.name, "lebih muda")
	}

	fmt.Println("==========")

	programmerSatu := model.Programmer{
		Name:     "Erwin",
		Job:      "Senior Backend",
		ExpYears: 6,
		Skills: []string{
			"C++", "Java", "C#", "HTML", "CSS",
			"JS", "PHP", "Laravel", "Python", "Flask",
			"Go", "Spring Boot",
			"MySQL", "PostgreSQL", "Redis", "MongoDB",
			"REST API", "gRPC", "Kafka", "Docker",
		},
		Company: model.Company{
			Name:     "Yourpay",
			Industry: "Fintech",
			Age:      5,
		},
	}
	fmt.Println("==========")

	programmerSatu.PrintDetail()
	programmerSatu.TambahExp(10)
	programmerSatu.AddExpYears(10)

	fmt.Println("==========")
	randomSkill := programmerSatu.GetRandomSkill()
	fmt.Println(randomSkill)
}
