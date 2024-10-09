package topic

import "fmt"

func PlayDefer() {
	fmt.Println("PlayDefer start")

	// defer punya internal stack
	defer fmt.Println("PlayDefer end") // run last

	defer func() {
		fmt.Println("defer anonymous function") //biasanya untuk close resources, seperti connection sql, message-queue, dkk
	}()
	defer fmt.Println("last is first") // close channel

}
