package main

import (
	"fmt"
	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
	"sync"
	"time"
)

func Task5Part1(n int) {
	ss := sync.WaitGroup{}

	for i := 1; i <= n; i++ {
		go func(i int) {
			ss.Add(1)
			// Do smth
			time.Sleep(1)
			fmt.Printf("%d goroutine done\n", i)
			defer ss.Done()
		}(i)
	}

	ss.Wait()
}

type Task5Part2Struct struct {
	sync.Mutex
	i int64
}

func (mdl *Task5Part2Struct) Task5Part2() {
	mdl.Lock()
	mdl.i++
	defer mdl.Unlock()
}

func main() {
	Task5Part1(10)

	ss := Task5Part2Struct{}
	ss.Task5Part2()

}
