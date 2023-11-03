package main

import (
	"github.com/felipedias-dev/gopportunities/config"
	"github.com/felipedias-dev/gopportunities/router"
)

var logger *config.Logger

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}
	router.Initialize()
}
