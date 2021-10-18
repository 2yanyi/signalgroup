# siggroup
基于系统信号量的多任务并行管理组

<br>

## Installing

```
go get github.com/matsuwin/siggroup
```

## Quick Start

```go
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
```
```
work_2 ...
work_1 ...
:shutdown
```

<br>

## X errcause

```go
package main

import (
	"github.com/matsuwin/siggroup/x/errcause"
	"github.com/pkg/errors"
	"io/ioutil"
)

func mkError() error {
	_, err := ioutil.ReadFile("xxx.txt")
	return errors.New(err.Error())
}

func main() {
	
	// 错误恢复 recover call errcause.Keep
	defer func() {
		if ei := recover(); ei != nil {
			errcause.Keep(ei)
		}
	}()

	// 模拟一个错误抛出调用
	if err := mkError(); err != nil {
		panic(err)
	}
}
```
