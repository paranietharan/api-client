package main

import (
	"api-client/data"
	"api-client/dto"
	"api-client/helper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := data.InitProducts(); err != nil {
		fmt.Println("Failed to sync products:", err)
		return
	}

	// Debug logs
	fmt.Println("Products loaded:", len(data.Products))
	fmt.Printf("First product: %+v\n", data.Products[0])

	router := gin.Default()
	router.POST("/order", SubmitOrder)
	router.Run(":8080")
}

func SubmitOrder(c *gin.Context) {
	var newOrder dto.OrderReq

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var payload dto.Order
	payload.ID = newOrder.ID
	for _, v := range newOrder.Items {
		i, err := data.FindByName(v.Name)
		if err != nil {
			fmt.Println("error occured at " + v.Name)
			c.JSON(http.StatusBadRequest, "error while processing product")
			return
		}
		payload.Items = append(payload.Items, i)
	}

	fmt.Println(payload)
	err := helper.SubmitOrder(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusCreated, "Ordr created")
}
