package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkMutex9010(b *testing.B) {
	var (
		counter int
		mutex   sync.Mutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mutex.Lock()

				temp := counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)

				counter = 1000
				mutex.Unlock()
			}
		})
	})
}

func BenchmarkMutex5050(b *testing.B) {
	var (
		counter int
		mutex   sync.Mutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mutex.Lock()

				temp := counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter

				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000

				mutex.Unlock()
			}
		})
	})
}

func BenchmarkMutex1090(b *testing.B) {
	var (
		counter int
		mutex   sync.Mutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mutex.Lock()

				temp := counter
				fmt.Println(temp)

				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000

				mutex.Unlock()
			}
		})
	})
}

func BenchmarkMutexRW9010(b *testing.B) {
	var (
		counter int
		mutex   sync.RWMutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mutex.Lock()

				temp := counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)

				counter = 1000
				mutex.Unlock()
			}
		})
	})
}

func BenchmarkMutexRW5050(b *testing.B) {
	var (
		counter int
		mutex   sync.RWMutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mutex.Lock()

				temp := counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter
				fmt.Println(temp)
				temp = counter

				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000

				mutex.Unlock()
			}
		})
	})
}

func BenchmarkMutexRW1090(b *testing.B) {
	var (
		counter int
		mutex   sync.RWMutex
	)
	b.Run(strconv.Itoa(runtime.GOMAXPROCS(0)), func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				mutex.Lock()

				temp := counter
				fmt.Println(temp)

				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000
				counter = 1000

				mutex.Unlock()
			}
		})
	})
}
