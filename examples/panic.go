package main

import (
	"fmt"
	"github.com/matsuwin/siggroup"
	"github.com/pkg/errors"
	"os"
	"time"
)

func main() {
	fmt.Println("program running.")

	siggroup.Async(func() (_ error) {
		fmt.Println("work_1 ...")
		time.Sleep(time.Second)
		if _, err := os.ReadFile("xxx.txt"); err != nil {
			panic(errors.New(err.Error()))
		}
		return
	})

	// wait.
	siggroup.Wait(func() {
		fmt.Println("program shutdown.")
	})
}
