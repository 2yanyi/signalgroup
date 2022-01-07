package signalgroup

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {

	// 添加异步任务 work_1，后台持续运行。
	Async(func() (_ error) {
		fmt.Println("work_1 ...")
		for {
		}
	})

	// 添加异步任务 work_2，短暂运行后退出。
	Async(func() (_ error) {
		fmt.Println("work_2 ...")
		time.Sleep(time.Second)
		return
	})

	// 等待任务结束，注意！只要有一个任务退出就退出所有。
	Wait(func() {
		fmt.Println(":shutdown")
		return
	})
}
