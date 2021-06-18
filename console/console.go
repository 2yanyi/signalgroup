////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// # 特性
// - 异步的日志记录，减小带来的程序阻塞
// - 支持屏幕打印和磁盘文件输出，文件输出是强制项
// - 支持磁盘文件滚动（滚动暂未支持）
// - 彩色的屏幕打印

package console

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"runtime"
	"sync"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const _INFO, _DEBUG, _WARN, _ERROR = 0, 1, 2, 3

type Options struct {
	Info     bool
	Debug    bool
	Warning  bool
	Error    bool
	Print    bool
	Filename string
}

type structure struct {
	L int    // Level
	F string // Function
	E error  // Error
}

var fos *os.File
var (
	control   = Options{true, true, true, true, true, ""}
	engine    = make(chan structure, 1024)
	wg        = sync.WaitGroup{}
	errorExit = errors.New("console: Exit")
)

// Wait 等待缓冲区的所有数据被消费
func (*structure) Wait() {
	engine <- structure{E: errorExit}
	wg.Wait()
}

// New 启用功能和配置输入，并等待结束
//
// main:
//   defer New(&Options{
//     Info: true, Debug: true, Warning: true, Error: true,
//     Print:    true,
//     Filename: "console.log",
//   }).Wait()
//
func New(options *Options) interface{ Wait() } {
	if options != nil {
		control = *options
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		loop()
	}()
	return &structure{}
}

func loop() {
	defer func() { _ = fos.Close() }()

	if control.Filename == "" {
		control.Filename = "console.log"
	}
	var err error
	fos, err = os.OpenFile(control.Filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	for elem := range engine {
		if elem.E == errorExit {
			return
		}
		now := time.Now().Local().Format(time.RFC3339)

		switch elem.L {
		case _INFO:
			_, _ = fos.WriteString(fmt.Sprintf("INFO   %s (%s)  %s\n", now, elem.F, elem.E))
		case _DEBUG:
			_, _ = fos.WriteString(fmt.Sprintf("DEBUG  %s (%s)  %s\n", now, elem.F, elem.E))
		case _WARN:
			_, _ = fos.WriteString(fmt.Sprintf("WARN   %s (%s)  %s\n", now, elem.F, elem.E))
		case _ERROR:
			_, _ = fos.WriteString(fmt.Sprintf("ERROR  %s (%s)  %s\n", now, elem.F, elem.E))
		}

		if control.Print {
			// fmt.Printf("\033[0;31;48m%s\033[0m\n", "RED")
			//
			// 前景 背景 颜色
			// ---------------------------------------
			// 30  40  黑色
			// 31  41  红色
			// 32  42  绿色
			// 33  43  黄色
			// 34  44  蓝色
			// 35  45  紫红色
			// 36  46  青蓝色
			// 37  47  白色
			//
			// 代码 意义
			// -------------------------
			//  0  终端默认设置
			//  1  高亮显示
			//  4  使用下划线
			//  5  闪烁
			//  7  反白显示
			//  8  不可见
			switch elem.L {
			case _INFO:
				fmt.Printf("INFO  %s (%s)  %s\n", now, elem.F, elem.E)
			case _DEBUG:
				fmt.Printf("\u001B[0;34;48mDEBUG %s (%s)  %s\u001B[0m\n", now, elem.F, elem.E)
			case _WARN:
				fmt.Printf("\u001B[0;33;48mWARN  %s (%s)  %s\u001B[0m\n", now, elem.F, elem.E)
			case _ERROR:
				fmt.Printf("\u001B[0;31;48mERROR %s (%s)  %s\u001B[0m\n", now, elem.F, elem.E)
			}
		}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// INFO 程序执行的过程信息
func INFO(msg string, a ...interface{}) {
	if control.Info {
		if len(a) != 0 {
			msg = fmt.Sprintf(msg, a...)
		}
		pc, _, _, _ := runtime.Caller(1)
		engine <- structure{_INFO, runtime.FuncForPC(pc).Name(), errors.New(msg)}
	}
}

// DEBUG 调试信息
func DEBUG(msg string, a ...interface{}) {
	if control.Debug {
		if len(a) != 0 {
			msg = fmt.Sprintf(msg, a...)
		}
		pc, _, _, _ := runtime.Caller(1)
		engine <- structure{_DEBUG, runtime.FuncForPC(pc).Name(), errors.New(msg)}
	}
}

// WARN 警告信息
func WARN(msg string, a ...interface{}) {
	if control.Warning {
		if len(a) != 0 {
			msg = fmt.Sprintf(msg, a...)
		}
		pc, _, _, _ := runtime.Caller(1)
		engine <- structure{_WARN, runtime.FuncForPC(pc).Name(), errors.New(msg)}
	}
}

// ERROR 错误信息
func ERROR(err error) {
	if control.Error {
		pc, _, _, _ := runtime.Caller(1)
		engine <- structure{_ERROR, runtime.FuncForPC(pc).Name(), err}
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// ManuallyClose 手动关闭日志
func ManuallyClose() {
	engine <- structure{E: errorExit}
}

// Timekeeper 函数运行计时
func Timekeeper(name string) func() {
	pc := make([]uintptr, 1)
	runtime.Callers(3, pc)
	start := time.Now()
	INFO("[ Timekeeper ] %s", name)
	return func() {
		INFO("[ Timekeeper ] %s T:%s", name, time.Since(start))
	}
}

// Json 格式化输出JSON
func Json(data interface{}, indent string) {
	if indent != "" {
		js, _ := json.MarshalIndent(data, "", indent)
		fmt.Println(string(js))
	} else {
		js, _ := json.Marshal(data)
		fmt.Println(string(js))
	}
}
