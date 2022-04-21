package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"fmt"
)

type uriBinding struct {
	ID string `uri:"id"`
}

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
}


type ProductForm struct {
	ID string `form:"id"`
	Name string `form:"name"`
}

type headerBinding struct {
	RequestID string `header:"X-Request-ID"`
}


func main() {
	var router = gin.Default()
	var address  = ":3000" 


	// c.ShouldBindUri

	router.GET("modelBinding/:id", func(c *gin.Context){
		
		var binding uriBinding
		
		if e := c.ShouldBindUri(&binding); e != nil {

			c.String(http.StatusBadRequest, e.Error())
			return

		}

		fmt.Println("binding is > ", binding)
		c.String(http.StatusOK, "Model Binding")

	})




	// c.ShouldBindJSON

	router.POST("productBinding/:id", func(c *gin.Context){
		
		var product Product
		
		if e := c.ShouldBindJSON(&product); e != nil {

			c.String(http.StatusBadRequest, e.Error())
			return

		}

		fmt.Println("product is > ", product)
		c.String(http.StatusOK, "Model Binding")

	})



	// c.ShouldBind     form data


	router.POST("modelFormBinding/:id", func(c *gin.Context){
		
		var productForm ProductForm
		
		if e := c.ShouldBind(&productForm); e != nil {

			c.String(http.StatusBadRequest, e.Error())
			return

		}

		fmt.Println("product is > ", productForm)
		c.String(http.StatusOK, "Model Binding")

	})




	// c.ShouldBindheader

	router.POST("modelHeaderBinding/:id", func(c *gin.Context){
		
		var product headerBinding
		
		if e := c.ShouldBindHeader(&product); e != nil {

			c.String(http.StatusBadRequest, e.Error())
			return

		}

		fmt.Println("product is > ", product)
		c.String(http.StatusOK, "Model Binding")

	})

	


 



	log.Fatalln(router.Run(address))

}

