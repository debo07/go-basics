package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _,link := range links {
		go checkLink(link, c)
	}
	// for i:=0; i <len(links); i++ {
	// 	fmt.Println("Index:", i)
	// 	fmt.Println(<- c) // This is a blocking call, loop will wait for the message to arrive on the channel, then it will move forward
	// }

	// This is the approach for infinite for loop, where checkLink gets called with same returned link over the channel
	// for {
	// 	go checkLink(<-c, c)
	// }

	// This is another way/syntx to check it in infinite loop, where l gets the msg sent over the channel c
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// Infinite loop but with a delay: Using function literal (annonymous function)
	for l:= range c {
		go func(link string) {
			time.Sleep(time.Second * 3) // sleep the child routine for 3 seconds before calling again.
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}