package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/klurpicolo/finalexam/models"
	"github.com/klurpicolo/finalexam/service"
)

func createCustomerHandler(c *gin.Context) {
	fmt.Println("createCustomerHandler")
	newCustomer := models.Customer{}
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		log.Println("ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCustomerID, err := service.Insert(&newCustomer)
	if err != nil {
		log.Println("Insert error", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	log.Println("Create success : Id=%v", createdCustomerID)
	newCustomer.ID = createdCustomerID
	c.JSON(http.StatusCreated, newCustomer)
}

func getAllCustomersHandler(c *gin.Context) {
	fmt.Println("getAllCustomersHandler")
	customers, err := service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customers)
}

func getCustomerByIdHandler(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("getCustomerByIdHandler")
	customer, err := service.FindbyID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/customers", createCustomerHandler)
	router.GET("/customers", getAllCustomersHandler)
	router.GET("/customers/:id", getCustomerByIdHandler)

	return router
}
