package main

import (
	"fmt"
	"github.com/matsuwin/siggroup"
	"time"
)

func main() {

	// 添加异步任务 work_1，后台持续运行。
	siggroup.Async(func() {
		fmt.Println("work_1 ...")
		for {}
	})

	// 添加异步任务 work_2，短暂运行后退出。
	siggroup.Async(func() {
		fmt.Println("work_2 ...")
		time.Sleep(time.Second)
	})

	// 等待任务结束，注意！只要有一个任务退出就退出所有。
	siggroup.Wait(func() {
		fmt.Println(":shutdown")
	})
}
