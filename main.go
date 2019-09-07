package main

import (
	"github.com/PlayList.API/apiRoots"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	api := app.Group("/kasim")

	apiRoots.Employee(api)

	app.Run("1111")
}
