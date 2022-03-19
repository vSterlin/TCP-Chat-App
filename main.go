package main

import (
	"github.com/vSterlin/chat/app"
)

func main() {

	// defer c.CloseConnections()

	app := app.NewApp()
	app.Run()

}
