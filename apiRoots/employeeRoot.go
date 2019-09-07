package apiRoots

import (
	"github.com/gin-gonic/gin"
)

func EmployeeAPI(api *gin.RouterGroup) {
	// CreateReadUpdateDelete = CRUD

	// Create employee model           (CREATE)
	api.POST("/employeePOST")         // localhost:1111/api//employeePOST
	// Read employee model             (READ)
	api.GET("/employeesGET")          // localhost:1111/api//employeeGET
	// Read id option employee model   (READ)
	api.GET("/employeesGET:id")       // localhost:1111/api//employeeGET:id
	// Update employee model           (UPDATE)
	api.PUT("/employeeUpdate")        // localhost:1111/api//employeeUpdate
	// Delete employee model           (DELETE)
	api.DELETE("/employeeDELETE")     // localhost:1111/api//employeeDELETE
}
