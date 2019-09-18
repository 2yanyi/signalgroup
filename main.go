package main

import (
	"github.com/zserge/lorca"
)

func main() {
	urlX := "https://im.dingtalk.com"
	ui, err := lorca.New(urlX, "", 1000, 600)
	if err != nil {
		panic(err)
	} else {
		defer ui.Close()
	}
	<-ui.Done()
}
