package main

import (
	"fmt"
	// "sync"
	"time"
)

func main() {
	// usingSleep()
	// usingScanln()
	// usingWaitGroup()
	usingChannel()
}

// func serve(personName string) {
// 	for i := 0; true; i++ {
// 		fmt.Println("serving customer", personName)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// func serveTemporary(personName string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("serving customer", personName)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

func serveWithChannel(personName string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("serving customer: %s", personName)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}


// func usingSleep() {
// 	go serve("husen") 
// 	go serve("pablo")
// 	go serve("armand")

// 	time.Sleep(time.Second * 4) // not useful, needs to manually enter sleep duration
// 	fmt.Println("Time's up guys. Restaurant wiil be closed")
// }


// func usingScanln() {
// 	go serve("husen") 
// 	go serve("pablo")
// 	go serve("armand")

// 	fmt.Scanln() // not useful way, because it needs user input (enter key)
// 	fmt.Println("Time's up guys. Restaurant wiil be closed")
// }


// func usingWaitGroup() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	func ()  {
// 		serveTemporary("fadhil")
// 		fmt.Println("Time's up guys. Restaurant wiil be closed")
// 		wg.Done()
// 	}()

// 	wg.Wait()
// }


func usingChannel() {
	c := make(chan string)
	go serveWithChannel("fadhil", c)

	// long version :
	// for {
	// 	message, isOpen := <- c
	// 	if !isOpen {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }
	
	//short version :
	for msg := range c {
		fmt.Println(msg)
	}
	fmt.Println("Time's up guys. Restaurant wiil be closed")
}