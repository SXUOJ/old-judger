package main

import (
	"github.com/isther/judger/web"
)

func main() {
	app := web.NewApp()
	app.Run()
}
