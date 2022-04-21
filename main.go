package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"fmt"
)


func main() {
	var router = gin.Default()
	var address  = ":3000" 


	// c.Param

	router.GET("product/:id", func(c *gin.Context){
		var id = c.Param("id")

		c.String(http.StatusOK, "ID is ", id)

	})


	// c.Query
	router.GET("productQuery/:id", func(c *gin.Context){
		var id = c.Query("id")
		var idDefaultValue = c.DefaultQuery("id_is", "12")

		fmt.Println("ID : ", id , "Default", idDefaultValue)

		c.String(http.StatusOK, "ID")

	})
	 

	// c.PostForm

	router.POST("productQuery/:id", func(c *gin.Context){
		var id = c.PostForm("id")
		var idDefaultValue = c.DefaultPostForm("id_is", "12")

		fmt.Println("ID : ", id , "Default  : ", idDefaultValue)

		c.String(http.StatusOK, "ID")

	})

	// c.GetHeader

	router.POST("productHeaderQuery/:id", func(c *gin.Context){
		var id = c.GetHeader("id")

		fmt.Println("ID : ", id )

		c.String(http.StatusOK, "ID")

	})



	log.Fatalln(router.Run(address))

}

