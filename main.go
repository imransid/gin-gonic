package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func main() {
	var router = gin.Default()
	var address  = ":3000" 


	var v1 = router.Group("/api/v1")

	v1.GET("/hello",  func(c *gin.Context){
		c.String(http.StatusOK, " Hi I M From V1")
	})

	var v2 = router.Group("/api/v2")

	v2.GET("/hello",  func(c *gin.Context){
		c.String(http.StatusOK, " Hi I M From V2")
	})

	router.GET(
		"/hello", func(c *gin.Context){
			c.String(http.StatusOK, "hello world ...")
		})

	log.Fatalln(router.Run(address))

}

