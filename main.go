package main

import (
	"os"

	"github.com/PedroBSanchez/gojobs.git/config"
	"github.com/PedroBSanchez/gojobs.git/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	print(os.Getenv("PORT"))

	//Initialize Configs
	err := config.Init()

	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()

}
