////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package hystrix

import (
	"errors"
	"github.com/matsuwin/fuseutil/siggroup/errcause"
	"time"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var ErrorTimeout = errors.New("hystrix: Timeout")

// Try 熔断降级工具
// err := Try(time.Second, func() error {
// 期望的正常处理
// ...
// return nil
// }, func(err error) error {
// 发生意外时的降级处理
// ...
// return err
// })
func Try(d time.Duration, run func() error, fallback func(e error) error) error {
	var sig = make(chan struct{})
	var err error

	// Begin
	go func() {
		defer func() {
			sig <- struct{}{}
			if ei := recover(); ei != nil {
				errcause.Keep(ei.(error))
			}
		}()
		err = run()
	}()

	// Wait
	select {
	case <-time.After(d):
		err = ErrorTimeout
	case <-sig:
	}

	// Fuse
	if fallback != nil && err != nil {
		err = fallback(err)
	}
	return err
}
