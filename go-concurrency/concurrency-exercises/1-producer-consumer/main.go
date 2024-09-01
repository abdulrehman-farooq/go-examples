//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

var channel chan *Tweet

func producer(stream Stream) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			close(channel)
			return
		}

		channel <- tweet
	}
}

func consumer() {
	for t := range channel {
		verifyTweets(*t)
	}
}

func verifyTweets(t Tweet) {

	if t.IsTalkingAboutGo() {
		fmt.Println(t.Username, "\ttweets about golang")
	} else {
		fmt.Println(t.Username, "\tdoes not tweet about golang")
	}
}
func main() {
	start := time.Now()
	stream := GetMockStream()
	channel = make(chan *Tweet)
	// Producer
	go producer(stream)

	// Consumer
	consumer()

	fmt.Printf("Process took %s\n", time.Since(start))
}
