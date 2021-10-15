package errcause_test

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"testing"

	"github.com/matsuwin/siggroup/x/errcause"
)

func Test(t *testing.T) {
	defer func() {
		if ei := recover(); ei != nil {
			errcause.Keep(ei)
		}
	}()

	if err := mkError(); err != nil {
		panic(err)
	}
}

func mkError() error {
	_, err := ioutil.ReadFile("xxx.txt")
	return errors.New(err.Error())
}
