package errcause_test

import (
	"github.com/matsuwin/siggroup/x/errcause"
	"github.com/pkg/errors"
	"io/ioutil"
	"testing"
)

func mkError() error {
	_, err := ioutil.ReadFile("xxx.txt")
	return errors.New(err.Error())
}

func Test(t *testing.T) {

	// 错误恢复 recover call errcause.Keep
	defer func() {
		if ei := recover(); ei != nil {
			errcause.Keep(ei)
		}
	}()

	// 模拟一个错误抛出调用
	if err := mkError(); err != nil {
		panic(err)
	}
}
