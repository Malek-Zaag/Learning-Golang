package main

import (
	"fmt"
	"sync"
)

type User struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var variable string = "hello world !!!"
	var username string
	fmt.Scan(&username)
	fmt.Println(variable + " " + username)
	bookings := []string{} //slice
	// var bookins [50] string array, you have to deal with index
	bookings = append(bookings, username, username)
	fmt.Println(bookings)
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	i := 3
	if i == 1 {
		fmt.Print(i)
	} else if i == 2 {
		fmt.Print(i)
	} else {
		fmt.Println("MMMm")
	}
	switch i {
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("three")
	}

	userData := make(map[string]string) //dictionnary
	var userTable []string
	//Goroutines => concurrent threads
	//creating threads is more expensie and slow startup time
	//heavyweight and needs more hardware resources(RAM)
	//"go ... " - starts a new goroutine
	//A goroutine is a lightweight thread managed by the Go runtime
	var wg = sync.WaitGroup{}
	wg.Add(1)
	//thread goes here
	//wg.done() removes thread from waiting list, should be in the function
	// called by the thread
	wg.Wait()
	//channels
	channel := make(chan string, 2)
	channel <- "hello"
	channel <- "hello"
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}
