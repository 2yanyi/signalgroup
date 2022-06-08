# signalgroup

Async work parallel controller based on system signals.

<br>

## Quick Start

```go
// Add work_1 in background.
signalgroup.Async(func() (_ error) {
    fmt.Println("work_1 ...")
    for {}
})

// Add work_2 keep 1s.
signalgroup.Async(func() (_ error) {
    fmt.Println("work_2 ...")
    time.Sleep(time.Second)
    return
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
