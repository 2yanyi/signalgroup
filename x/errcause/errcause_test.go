package errcause_test

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"testing"

	"github.com/matsuwin/siggroup/x/errcause"
)

func Test(t *testing.T) {
	defer errcause.Recover()

	if err := mkError(); err != nil {
		panic(err)
	}
}

func mkError() (_ error) {
	_, err := ioutil.ReadFile("xxx.txt")
	if err != nil {
		return errors.New(err.Error())
	}
	return
}
