# signalgroup

Async work parallel controller based on system signals.

<br>

## Quick Start

```go
signalgroup.Async(func() error {
    fmt.Println("work_1 ...")
    for {}
})

signalgroup.Async(func() error {
    fmt.Println("work_2 ...")
    time.Sleep(time.Second)
    return nil
})

// Wait end. If one work exits, it ends all.
signalgroup.Wait(func() {
    fmt.Println("<- shutdown")
    return
})
```
```
work_2 ...
work_1 ...
<- shutdown
```
