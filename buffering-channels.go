package main

import "fmt"

func main() {
	myArray := []int{1, 2}

	//Creating a buffered channel with the size of the array
	messages := make(chan string, len(myArray))

	for i := 0; i < len(myArray); i++ {
		//Sending a message through channel
		messages <- fmt.Sprintf("buffered channel: %d", myArray[i])
	}

	//Sending another message through channel
	//messages <- "chanell"

	//Printing output of channel with <- operator
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
