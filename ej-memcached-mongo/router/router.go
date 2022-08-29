package router

import (
	"fmt"
	bookController "github.com/emikohmann/ucc-arqsoft-2/ej-memcached-mongo/controllers/book"
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {
	// Products Mapping
	router.GET("/books/:id", bookController.Get)
	router.POST("/books", bookController.Insert)

	fmt.Println("Finishing mappings configurations")
}
