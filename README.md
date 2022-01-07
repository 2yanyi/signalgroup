# signalgroup

基于操作系统信号量的异步任务并行控制器，思想借鉴 errgroup。

<br>

## Quick Start

```go
// 添加异步任务 work_1，后台持续运行。
signalgroup.Async(func() (_ error) {
    fmt.Println("work_1 ...")
    for {}
})

// 添加异步任务 work_2，短暂运行后退出。
signalgroup.Async(func() (_ error) {
    fmt.Println("work_2 ...")
    time.Sleep(time.Second)
    return
})

// 等待任务结束，注意！只要有一个任务退出就退出所有。
signalgroup.Wait(func() {
    fmt.Println(":shutdown")
    return
})
```
```
work_2 ...
work_1 ...
:shutdown
```

## Installing

```
go get github.com/matsuwin/signalgroup
```
