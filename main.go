package main

import (
	"api-campaign/router"
	"api-campaign/utils"
)

func main() {
	route := router.SetupRouter(utils.GetConfig("GIN_MODE"))

	route.Run(":" + utils.GetConfig("PORT"))
}