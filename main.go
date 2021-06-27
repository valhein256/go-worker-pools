package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func sleep() {
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Millisecond,
	)
}

type Item struct {
	name  string
	value int
}

func echoWorker(in, out chan *Item) {
	for {
		item := <-in

		item.name = fmt.Sprintf("%s-After-worker", item.name)

		out <- item
	}
}

func producter(ch chan<- *Item) {
	i := 0
	for {
		item := &Item{
			name:  fmt.Sprintf("%s-%d", "P", i),
			value: i,
		}
		fmt.Printf("Send %v\n", item)
		ch <- item
		i++
	}
}

func main() {

	var num int
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			num = 3
		} else {
			num = i
		}
	} else {
		num = 3
	}
	println(num)
	in := make(chan *Item)
	out := make(chan *Item)
	for i := 1; i <= num; i++ {
		go echoWorker(in, out)
	}
	go producter(in)
	for item := range out {
		fmt.Printf("Recv %v from channel.\n", item)
	}
}
