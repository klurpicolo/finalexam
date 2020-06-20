package routers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/klurpicolo/finalexam/models"
	"github.com/klurpicolo/finalexam/service"
)

func createCustomerHandler(c *gin.Context) {
	log.Println("createCustomerHandler")
	newCustomer := models.Customer{}
	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		log.Println("ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCustomerID, err := service.Insert(&newCustomer)
	if err != nil {
		log.Println("Insert error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Create success : Id= ", createdCustomerID)
	updatedID, err := strconv.Atoi(createdCustomerID)
	if err != nil {
		log.Println("Convert error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newCustomer.ID = updatedID
	c.JSON(http.StatusCreated, newCustomer)
}

func getAllCustomersHandler(c *gin.Context) {
	log.Println("getAllCustomersHandler")
	customers, err := service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}

func getCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	log.Println("getCustomerByIdHandler")
	customer, err := service.FindbyID(id)
	if err != nil {
		log.Println("getCustomerByIdHandler error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}

func updateCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	log.Println("updateCustomerByIdHandler")

	updatedCustomer := models.Customer{}
	if err := c.ShouldBindJSON(&updatedCustomer); err != nil {
		log.Println("ShouldBindJSON", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check if customer exists
	if _, err := service.FindbyID(id); err != nil {
		log.Println("Customer not exist", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateByID(id, &updatedCustomer); err != nil {
		log.Println("Update error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Update success : Id= ", id)
	c.JSON(http.StatusOK, updatedCustomer)
}

func deleteCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	log.Println("deleteCustomerByIDHandler")

	//Check if customer exists
	if _, err := service.FindbyID(id); err != nil {
		log.Println("Customer not exist", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.DeleteByID(id); err != nil {
		log.Println("Delete error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Delete success : Id= ", id)
	c.JSON(http.StatusOK, gin.H{
		"message": "customer deleted",
	})
}

func authMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "token2019" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	c.Next()
}

//GetRouter comment
func GetRouter() *gin.Engine {
	router := gin.Default()

	router.Use(authMiddleware)

	router.POST("/customers", createCustomerHandler)
	router.GET("/customers", getAllCustomersHandler)
	router.GET("/customers/:id", getCustomerByIDHandler)
	router.PUT("/customers/:id", updateCustomerByIDHandler)
	router.DELETE("/customers/:id", deleteCustomerByIDHandler)

	return router
}
