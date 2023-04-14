package main

import "fmt"

// func UsingSleep() {
// 	go Serve("husen")
// 	go Serve("pablo")
// 	go Serve("armand")

// 	time.Sleep(time.Second * 4) // not useful, needs to manually enter sleep duration
// 	fmt.Println("Time's up guys. Restaurant wiil be closed")
// }

// func UsingScanln() {
// 	go Serve("husen")
// 	go Serve("pablo")
// 	go Serve("armand")

// 	fmt.Scanln() // not useful way, because it needs user input (enter key)
// 	fmt.Println("Time's up guys. Restaurant wiil be closed")
// }

// func UsingWaitGroup() {
// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	func ()  {
// 		ServeTemporary("fadhil")
// 		fmt.Println("Time's up guys. Restaurant wiil be closed")
// 		wg.Done()
// 	}()

// 	wg.Wait()
// }


func UsingChannel() {
	c := make(chan string)
	go ServeWithChannel("fadhil", c)

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