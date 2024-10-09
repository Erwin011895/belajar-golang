package model

import (
	"fmt"
	"math/rand"
)

type Programmer struct {
	Name     string
	Job      string
	ExpYears int
	Skills   []string
	Company  Company
}

func (p Programmer) PrintDetail() {
	fmt.Println("nama saya", p.Name)
	fmt.Println("saat ini saya bekerja sebagai", p.Job)
	fmt.Println("saya sudah punya pengalaman", p.ExpYears, "tahun")

	fmt.Println("skill saya:")
	for i, skill := range p.Skills {
		fmt.Println(i+1, skill)
	}

	fmt.Println("saat ini saya bekerja di", p.Company.Name, "yang bergerak di bidang", p.Company.Industry)
}

// pointer receiver. CAN mutate p values
func (p *Programmer) AddExpYears(x int) {
	p.ExpYears += x
}

// non-pointer receiver. CANNOT mutate p values
func (p Programmer) TambahExp(x int) {
	p.ExpYears += x // ada warning: ineffective assignment to field Programmer.ExpYears
}

func (p Programmer) GetRandomSkill() string {
	randomIndex := rand.Intn(len(p.Skills))
	return p.Skills[randomIndex]
}
