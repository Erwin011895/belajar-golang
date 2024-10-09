package topic

import (
	"fmt"
	"time"
)

func BasicChan() {
	/* UNBUFFERED CHANNEL */
	messageChan := make(chan string) // unbuffered channel, bisa terima unlimited message
	go func() {
		messageChan <- "tes" // send "tes" to messageChan
	}()
	msg := <-messageChan // output 1 data from messageChan. possible BLOCKING when no data in messageChan
	fmt.Println("msg", msg)
	// msg = <-messageChan // uncomment this will BLOCK main goroutine, waiting for new data from messageChan
	// return

	/* BUFFERED CHANNEL */
	bufferedMessageChan := make(chan string, 2) // maksimal 2 message di dalam chan dalam 1 waktu

	go func() {
		for i := 1; i <= 10; i++ {
			bufferedMessageChan <- fmt.Sprintf("data %d", i)
			fmt.Println("done sending data", i)
			// time.Sleep(time.Second)
		}
		close(bufferedMessageChan)
	}()

	// range over channel. quit loop when channel is CLOSED
	for str := range bufferedMessageChan {
		fmt.Println("received", str)
		time.Sleep(time.Second)
	}

	/* CHANNEL DIRECTION in func */
	// data angka 2 => chPertama => chKedua (convert jd text) => output di main gorutine
	chPertama := make(chan int, 1)
	chKedua := make(chan string, 1)
	passMsg1(chPertama, 2)
	passMsg2(chPertama, chKedua)

	fmt.Println("chKedua:", <-chKedua)
}

func passMsg1(ch chan<- int, data int) {
	fmt.Println("passMsg1", data)
	ch <- data // accepted
	// <- ch // error
}

func passMsg2(ch1 <-chan int, ch2 chan<- string) {
	fmt.Println("passMsg2")
	angka := <-ch1
	texts := []string{"nol", "satu", "dua", "tiga", "empat"}
	ch2 <- texts[angka]
	// ch1 <- 2 // error
}
