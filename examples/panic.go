package main

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"github.com/matsuwin/siggroup"
	"time"
)

func main() {
	fmt.Println("program running.")

	siggroup.Async(func() {
		fmt.Println("work_1 ...")
		time.Sleep(time.Second)
		if _, err := ioutil.ReadFile("xxx.txt"); err != nil {
			panic(errors.New(err.Error()))
		}
	})

	// wait.
	siggroup.Wait(func() {
		fmt.Println("program shutdown.")
	})
}
