package main

import (
	"fmt"
	"time"
)

// func Serve(personName string) {
// 	for i := 0; true; i++ {
// 		fmt.Println("serving customer", personName)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func ServeTemporary(personName string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("serving customer", personName)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

func ServeWithChannel(personName string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("serving customer: %s", personName)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}