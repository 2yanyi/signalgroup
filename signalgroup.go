package signalgroup

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/matsuwin/errcause"
)

func exitHistory(message string) {
	fp := "exit.history"
	ti := time.Now().Local().Format("20060102.150405")
	sh := fmt.Sprintf("echo '%s %s' >> %s", ti, message, fp)
	var err error
	if runtime.GOOS == "windows" {
		err = exec.Command("cmd", "/c", sh).Run()
	} else {
		err = exec.Command("sh", "-c", sh).Run()
	}
	if err != nil {
		println(err.Error())
	}
}

// Quit Send process exit signal
func Quit() {
	sig <- syscall.SIGQUIT
}

// Async new Goroutine
func Async(worker func() error) {
	atomic.AddInt32(&countWork, 1)
	go func() {
		defer errcause.Recover()
		defer Quit()
		_ = worker()
	}()
}

// Wait Listen process exit signal
func Wait(cancel func()) {
	if countWork == 0 {
		return
	}
	ls := make([]os.Signal, 0, 15)
	for i := range signals {
		ls = append(ls, i)
	}
	signal.Notify(sig, ls...)
	for message := range sig {
		exitHistory(signals[message])
		if cancel != nil {
			cancel()
		}
		return
	}
}

var signals = map[os.Signal]string{
	syscall.SIGHUP:  "(SIGHUP)  1  终端挂断，终端 session 结束。",
	syscall.SIGINT:  "(SIGINT)  2  进程中断，如 Ctrl-C，通常由用户触发。",
	syscall.SIGQUIT: "(SIGQUIT) 3  进程退出，如 Ctrl-\\，或内部 Quit 调用。",
	syscall.SIGILL:  "(SIGILL)  4  非法指令，代码错误。",
	syscall.SIGFPE:  "(SIGFPE)  8  运算异常，代码错误。",
	syscall.SIGSEGV: "(SIGSEGV) 11 分段冲突，代码错误。",
	syscall.SIGPIPE: "(SIGPIPE) 13 管道破裂，进程间通信故障。",
	syscall.SIGTERM: "(SIGTERM) 15 终止进程，如 kill，通常由进程外部触发。",
}

var sig = make(chan os.Signal)
var countWork int32
