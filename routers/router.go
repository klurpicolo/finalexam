package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klurpicolo/finalexam/service"
)

func getAllCustomersHandler(c *gin.Context) {
	fmt.Println("getAllCustomersHandler")
	customers, err := service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	// customers = append(customers, &Customer{"1", "klur", "lkur@giam.com", "active"})

	c.JSON(http.StatusOK, customers)
}

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/customers", getAllCustomersHandler)

	return router
}
