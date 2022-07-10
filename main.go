package main

import (
	"github.com/iloncar89/banking-lib/logger"
	"github.com/iloncar89/banking/app"
)

func main() {

	logger.Info("Starting the application")
	app.Start()

}
