package main

import (
	"fmt"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()

	// wg.Wait()

	c := make(chan string)
	go count("sheep", c)

	// msg := <- c
	// fmt.Println(msg)

	// for {
	// 	msg, open := <- c
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(msg)
	// }

	for msg := range c {
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- fmt.Sprintf("%d %s", i, thing)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}