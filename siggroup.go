package siggroup

import (
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	"github.com/matsuwin/siggroup/x/errcause"
)

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
	signal.Notify(sig, _signal...)
	for message := range sig {
		for i := range _signal {
			if message == _signal[i] {
				_ = os.WriteFile("signal.txt", []byte(message.String()), 0666)
				if cancel != nil {
					cancel()
				}
				return
			}
		}
	}
}

var _signal = []os.Signal{
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
