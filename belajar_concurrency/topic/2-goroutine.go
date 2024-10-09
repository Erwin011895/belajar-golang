package topic

import (
	"fmt"
	"time"
)

func BasicGoroutine() {
	fmt.Println("basicGoroutine start")
	// spawn 1 goroutine with "anonymous func" / Closure
	go func() {
		fmt.Println("anonymous func goroutine")
	}() // JS IIFE

	// spawn 1 goroutine that run fmt.Println()
	go fmt.Println("hello")
	go ngePrint()

	// spawn 10 goroutine that each goroutine will run fmt.Println() 0 to 9 respectively
	for i := 0; i < 10; i++ {
		go fmt.Println(i)
	}
	time.Sleep(time.Second * 1) // blocking main-goroutine 1 sec, to make sure the goroutines spawned here running
}

func ngePrint() {
	fmt.Println("ngePrint")
}
