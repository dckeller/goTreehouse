package main 

import ("fmt"
				"math/rand"
				"time")


// This is writing to a channel
func longTask(channel chan int) {
	delay := rand.Intn(5)
	fmt.Println("Starting task")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("Ending task")
	channel <- delay 
}

// This is reading from the channel above
func main() {
	rand.Seed(time.Now().Unix())
	channel := make(chan int)

	for i := 1; i <= 3; i++ {
		go longTask(channel)
	}
	for i := 1; i <= 3; i++ {
		fmt.Println("Took", <-channel, "Seconds")
	}					
}

