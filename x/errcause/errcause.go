package errcause

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"reflect"
	"strings"
	"time"
)

// Fetch 从 error 中获取包含堆栈记录的错误根本原因
func Fetch(err error) string {
	message := fmt.Sprintf("panic: %+v", errors.Cause(err))
	if strings.Count(message, "runtime.goexit") == 0 {
		message = fmt.Sprintf("(Not github.com/pkg/errors.New) panic: %+v", errors.New(err.Error()))
	}
	return message
}

// Keep 内部调用 Fetch 并将结果输出到文件
func Keep(err interface{}) {
	RFC3339Nano := time.Now().Local().Format(time.RFC3339Nano)
	defer func() {
		/*
		 * 错误恢复 recover 工作方式
		 *
		 * 没有 recover:
		 *     触发 panic 后开始向上 (函数调用链) 传递错误，到达当前 goroutine 顶层时会退出整个进程！！！
		 *
		 * 有 recover:
		 *     触发 panic 后开始向上传递错误，遇见第一个 recover 后结束传递，达到恢复的效果。
		 */
		if ei := recover(); ei != nil {
			message := fmt.Sprintf("errcause: Error message keep failed: %s", ei)
			save(RFC3339Nano, message)
		}
	}()
	if reflect.TypeOf(err).String() == "string" {
		message := fmt.Sprintf("[  ERROR  ] %s -> %s\n", RFC3339Nano, err)
		save(RFC3339Nano, message)
	} else {
		message := Fetch(err.(error))
		message = fmt.Sprintf("[  ERROR  ] %s\n%s\n\n", RFC3339Nano, message)
		save(RFC3339Nano, message)
	}
}

func save(RFC3339Nano, message string) {
	fmt.Println(message)
	f, _ := os.OpenFile("panic."+RFC3339Nano[:10]+".log", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	_, _ = f.WriteString(message)
	_ = f.Close()
}
