package main

import (
	"github.com/vSterlin/tcp/app"
)

func main() {

	// defer c.CloseConnections()

	app := app.NewApp()
	app.Run()

}
