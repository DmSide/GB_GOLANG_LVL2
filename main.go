package main

import (
	"context"
	"fmt"
	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type MyLock struct {
	sync.Locker
}

type PP struct {
	sync.Mutex
}

func (s *PP) IncValue(ctx context.Context) {
	s.Lock()
	defer s.Unlock()
	// ctx.Value()
}

func Context() {
	p := PP{}
	ctx := context.WithValue(context.Background(), "val", 1)
	for i := 1; i <= 1000; i++ {
		go func() {
			p.IncValue(ctx)
		}()
	}
}

func AtomicTest() {
	c := int32(0)
	n := 1000
	ch := make(chan struct{}, n)

	go func() {
		for range ch {
			atomic.AddInt32(&c, 1)
		}
	}()

	for i := 0; i < n; i++ {
		//go func(i int) {
		ch <- struct{}{}
		//}(i)
	}

	fmt.Println(c)
}

func LectionExample() {
	var workers = make(chan struct{}, 1)
	var ch = make(chan int, 1)
	ch <- 1
	for i := 1; i <= 1000; i++ {
		workers <- struct{}{}

		go func() {
			defer func() {
				<-workers
				ch <- (<-ch) + 1
			}()
		}()
	}

	fmt.Println(<-ch)
}

func f1(ctx context.Context, sigs chan os.Signal) {

	for {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Printf("DeadlineExceed")
			}
			return
		case <-sigs:
			fmt.Printf("SIGTERM detected")
		default:
		}
	}
}

func Signals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	ctx := context.Background()
	cancelableCtx, _ := context.WithCancel(ctx)
	deadlineCtx, deadlineFunc := context.WithDeadline(cancelableCtx, time.Now().Add(time.Second))

	go f1(deadlineCtx, sigs)

	time.Sleep(5 * time.Second)

	defer deadlineFunc()
}

func Task4() {
	LectionExample()

	Signals()
}

func main() {
	Task4()
}
