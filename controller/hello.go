package controller

import (
	"fmt"
	"gin_api/model"
	"github.com/gin-gonic/gin"
)

// SayHello godoc
// @Summary      Say hello
// @Description  Says hello to a specified name
// @Tags         conversation
// @Param        HelloRequest  body  model.HelloRequest  true  "name to say hello to"
// @Produce      json
// @Router       /hello [post]
func SayHello(c *gin.Context) {
	requestBody := model.HelloRequest{}

	_ = c.ShouldBindJSON(&requestBody)
	fmt.Println("Sending response")
	c.JSON(200, gin.H{
		"message": "Hello " + requestBody.Name,
	})
	fmt.Println("Finishing request")
}
