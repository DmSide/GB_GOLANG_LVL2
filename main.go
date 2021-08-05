package main

import (
	"fmt"
	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func Race() {
	m := make(map[string]string)
	go func() {
		m["1"] = "a"
	}()
	m["2"] = "b"
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func MutexTrace() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	Mutex()
}

func Mutex() {
	const count = 1000

	var (
		counter int
		mutex   sync.Mutex

		// Вспомогательная часть нашего кода
		ch = make(chan struct{}, count)
	)
	for i := 0; i < count; i += 1 {
		go func() {
			// Захват мьютекса
			mutex.Lock()
			counter += 1
			// Освобождение мьютекса
			mutex.Unlock()

			// Фиксация факта запуска горутины в канале
			ch <- struct{}{}
		}()
	}
	time.Sleep(2 * time.Second)
	close(ch)

	i := 0
	for range ch {
		i += 1
	}
	// Выводим показание счетчика
	fmt.Println(counter)
	// Выводим показания канала
	fmt.Println(i)
}

func Trace() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	wg := sync.WaitGroup{}
	for i := 0; i < 1<<4; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1e8; i += 1 {
			}
			runtime.Gosched()
		}()
	}
	wg.Wait()
}

func main() {
	// Race()
	// Trace()
	MutexTrace()
}
