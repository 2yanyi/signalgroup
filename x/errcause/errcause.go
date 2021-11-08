package errcause

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"strings"
	"time"
)

/*
 * # 错误恢复: (recover) 的工作方式
 *
 * 没有 recover:
 *     触发 panic 后开始向上 (函数调用链) 传递错误，到达当前 goroutine 顶层时会退出整个进程！！！
 *
 * 有 recover:
 *     触发 panic 后开始向上传递错误，遇见第一个 recover 后结束传递，达到恢复的效果。
 */

// Cause 从 error 中获取包含堆栈记录的错误根本原因
func Cause(err error) string {
	message := fmt.Sprintf("panic: %+v", errors.Cause(err))
	if strings.Count(message, "runtime.goexit") == 0 {
		message = fmt.Sprintf("(Not github.com/pkg/errors.New) panic: %+v", errors.New(err.Error()))
	}
	return message
}

// Recover panic! Error recovery
//
// Old:
// go func() {
//     defer func() {
//         if ei := recover(); ei != nil {
//             // ...
//         }
//     }()
// }()
//
// New:
// go func() {
//     defer errcause.Recover()
// }()
//
func Recover() {
	if err := recover(); err != nil {
		RFC3339Nano := time.Now().Local().Format(time.RFC3339Nano)
		if witch {
			if reflect.TypeOf(err).String() == "string" {
				message := fmt.Sprintf("[  ERROR  ] %s -> %s\n", RFC3339Nano, err)
				save(RFC3339Nano, message)
			} else {
				message := fmt.Sprintf("[  ERROR  ] %s\n%s\n\n", RFC3339Nano, Cause(err.(error)))
				save(RFC3339Nano, message)
			}
			go func() {
				time.Sleep(time.Second)
				witch = true
			}()
		}
		witch = false
	}
}

func save(RFC3339Nano, message string) {
	fmt.Println(message)
	f, _ := os.OpenFile("panic."+RFC3339Nano[:10]+".log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	_, _ = f.WriteString(message)
	_ = f.Close()
}

var witch = true
