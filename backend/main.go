package main

import (
	"github.com/tedoham/xm-test/app"
	"github.com/tedoham/xm-test/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
