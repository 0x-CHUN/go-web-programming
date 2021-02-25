package main

import "fmt"

func callerA(c chan string) {
	c <- "Hello world!"
	close(c)
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	for i := 0; i < 5; i++ {
		select {
		case msg, ok1 := <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
		case msg, ok2 := <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
		}
	}
}
