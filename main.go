package main

import (
	"fmt"
	"math/rand"
	"mqtt-publisher/setup"
	"mqtt-publisher/topics"
	"time"
)

func main() {
	client, err := setup.Setup()
	if err != nil {
		panic(err)
	}

	go func() {
		num := 100
		allTopics := topics.AllTopics
		for i := 0; i <= num; i++ {
			rand.Seed(time.Now().UnixNano())
			randomInt := rand.Intn(2)
			topic := allTopics[randomInt]
			token := client.Publish(topic, 1, false, fmt.Sprintf("%s - %d - message payload", topic, randomInt))
			token.Wait()
			time.Sleep(time.Duration(randomInt) * time.Second)
		}
	}()

	<-make(chan int)
}
