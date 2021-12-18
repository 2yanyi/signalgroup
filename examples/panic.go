package main

import (
	"fmt"
	"github.com/matsuwin/siggroup"
	"github.com/pkg/errors"
	"io/ioutil"
	"time"
)

func main() {
	fmt.Println("program running.")

	siggroup.Async(func() (_ error) {
		fmt.Println("work_1 ...")
		time.Sleep(time.Second)
		if _, err := ioutil.ReadFile("xxx.txt"); err != nil {
			panic(errors.New(err.Error()))
		}
		return
	})

	// wait.
	siggroup.Wait(func() {
		fmt.Println("program shutdown.")
	})
}
