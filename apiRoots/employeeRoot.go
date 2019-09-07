package apiRoots

import (
	"github.com/gin-gonic/gin"
)

func EmployeeAPI(api *gin.RouterGroup) {
	// CreateReadUpdateDelete = CRUD

	// Create employee model           (CREATE)
	api.POST("/employeePOST")
	// Read employee model             (READ)
	api.GET("/employeesGET")
	// Read id option employee model   (READ)
	api.GET("/employeesGET:id")
	// Update employee model           (UPDATE)
	api.PUT("/employeeUpdate")
	// Delete employee model           (DELETE)
	api.DELETE("/employeeDELETE")
}
