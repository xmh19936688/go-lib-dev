package main

import (
	"fmt"
	"sync"

	"github.com/xmh19936688/go-dispatcher/dispatcher"
)

func main() {
	fmt.Println("go lib example")

	ch := make(chan string)
	wg := &sync.WaitGroup{}
	d := dispatcher.New[string]().Chan(ch).WaitGroup(wg).Handler(handler)
	go d.Start()

	data := []string{"1", "2"}
	wg.Add(len(data))
	for _, str := range data {
		ch <- str
	}

	wg.Wait()
}

func handler(arg string) {
	fmt.Println(arg)
}
