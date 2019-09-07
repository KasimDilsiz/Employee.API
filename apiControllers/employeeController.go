package apiControllers

import (
	"github.com/Employee.API/dbConfig"
	"github.com/Employee.API/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostEmployee(c *gin.Context) {

}

func GetEmployee(c *gin.Context) {
	var employees []models.Employee
	err := dbConfig.DB.Find(&employees).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, "Hatalı gönderim yaptınız. Hata: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, &employees)
}

func GetIdEmployee(c *gin.Context) {

}

func PutEmployee(c *gin.Context) {

}

func DeleteEmployee(c *gin.Context) {

}
