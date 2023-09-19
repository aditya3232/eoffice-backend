package main

import (
	"eoffice-backend/config"
	"eoffice-backend/helper"
	"eoffice-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// panic recovery
	defer helper.RecoverPanic()

	router := gin.Default()
	if config.ENV.DEBUG == 0 {
		gin.SetMode(gin.ReleaseMode)
	}

	routes.Initialize(router)
	router.Run()
}
