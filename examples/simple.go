package main

import (
	"fmt"
	"r/siggroup"
	"time"
)

func main() {
	fmt.Println("program running.")

	siggroup.Async(func() {
		fmt.Println("work_1 ...")
		for {
		}
	})

	siggroup.Async(func() {
		fmt.Println("work_2 ...")
		time.Sleep(time.Second)
	})

	// wait.
	siggroup.Wait(func() {
		fmt.Println("program shutdown.")
	})
}
