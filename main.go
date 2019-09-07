package main

import (
	"github.com/Employee.API/apiRoots"
	"github.com/Employee.API/dbConfig"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	dbConfig.InitDB()

	api := app.Group("/api")

	apiRoots.EmployeeAPI(api)

	app.Run(":1111")
}
