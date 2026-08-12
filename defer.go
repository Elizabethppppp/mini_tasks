package main

import (
	"fmt"
	"time"
)

func process1() {

	for i := 0; i < 3; i++ {

		fmt.Println("открываю", i)

		fmt.Println("работаю с", i)
		func() { defer fmt.Println("закрываю", i) }()
	}
}

func timed(name string) func() {
	start := time.Now()
	return func() { fmt.Printf("%s: %v\n", name, time.Since(start)) }
}
