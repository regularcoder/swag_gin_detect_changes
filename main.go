package main

import (
	"gin_api/controller"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"       // swagger embed files
import _ "gin_api/docs"

// @title          Example API
// @version      1.0
// @description  This is a sample server.

// @host  localhost:8080
func main() {
	r := gin.Default()

	r.POST("/hello", controller.SayHello)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
