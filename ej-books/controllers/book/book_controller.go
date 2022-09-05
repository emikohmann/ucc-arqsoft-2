package book

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/dtos"
	service "github.com/emikohmann/ucc-arqsoft-2/ej-books/services"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/services/repositories"
	"github.com/emikohmann/ucc-arqsoft-2/ej-books/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	serv = service.NewServiceImpl(
		repositories.NewRepositoryCache(),
		repositories.NewRepositoryMemcached(),
		repositories.NewRepositoryMongo(),
	)
)

func Get(c *gin.Context) {
	book, apiErr := serv.Get(c.Param("id"))
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusOK, book)
}

func Insert(c *gin.Context) {
	var book dtos.BookDto
	if err := c.BindJSON(&book); err != nil {
		apiErr := errors.NewBadRequestApiError(err.Error())
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	book, apiErr := serv.Insert(book)
	if apiErr != nil {
		c.JSON(apiErr.Status(), apiErr)
		return
	}
	c.JSON(http.StatusCreated, book)
}
