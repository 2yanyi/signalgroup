package main

import (
	"fmt"
	"siggroup"
	"time"
)

func main() {
	fmt.Println("program running.")

	siggroup.Add(func() {
		fmt.Println("work_1 ...")
		for {
		}
	})

	siggroup.Add(func() {
		fmt.Println("work_2 ...")
		time.Sleep(time.Second)
	})

	// wait.
	siggroup.Wait(func() {
		fmt.Println("program shutdown.")
	})
}
