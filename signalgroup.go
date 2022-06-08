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

func shutdownHistory(message string) {
	fp := "shutdown.history"
	ti := time.Now().Local().Format("2006-01-02.15:04:05")
	sh := fmt.Sprintf("echo %s %s >> %s", ti, message, fp)
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
func Async(routine func() error) {
	atomic.AddInt32(&countWork, 1)
	go func() {
		defer errcause.Recover()
		defer Quit()
		_ = routine()
	}()
}

// Wait Listen process exit signal
func Wait(cancel func()) {
	if countWork == 0 {
		return
	}
	signal.Notify(sig, signals...)
	for message := range sig {
		for i := range signals {
			if message == signals[i] {
				shutdownHistory(message.String())
				if cancel != nil {
					cancel()
				}
				return
			}
		}
	}
}

var signals = []os.Signal{
	syscall.SIGHUP,  //  1:  hangup
	syscall.SIGINT,  //  2:  interrupt
	syscall.SIGQUIT, //  3:  quit
	syscall.SIGILL,  //  4:  illegal instruction
	syscall.SIGTRAP, //  5:  trace/breakpoint trap
	syscall.SIGABRT, //  6:  aborted
	syscall.SIGBUS,  //  7:  bus error
	syscall.SIGFPE,  //  8:  floating point exception
	syscall.SIGSEGV, // 11:  segmentation fault
	syscall.SIGALRM, // 14:  alarm clock
	syscall.SIGTERM, // 15:  terminated
}

var sig = make(chan os.Signal)
var countWork int32
