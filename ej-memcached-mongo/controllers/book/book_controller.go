package book

import (
	"encoding/json"
	"fmt"
	"github.com/emikohmann/ucc-arqsoft-2/ej-memcached-mongo/dtos"
	service "github.com/emikohmann/ucc-arqsoft-2/ej-memcached-mongo/services"
	"github.com/emikohmann/ucc-arqsoft-2/ej-memcached-mongo/utils/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {

	id := c.Param("id")

	res := cache.Get(id)

	if res != "" {
		var bookDtoCache dtos.BookDto
		json.Unmarshal([]byte(res), &bookDtoCache)
		fmt.Println("from cache: " + id)
		c.JSON(http.StatusOK, bookDtoCache)
		return

	}
	bookDto, er := service.BookService.GetBook(id)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusOK, bookDto)
}

func Insert(c *gin.Context) {
	var bookDto dtos.BookDto
	err := c.BindJSON(&bookDto)

	// Error Parsing json param
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bookDto, er := service.BookService.InsertBook(bookDto)

	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	bookDtoStr, _ := json.Marshal(bookDto)
	cache.Set(bookDto.Id, bookDtoStr)
	fmt.Println("save cache: " + bookDto.Id)
	c.JSON(http.StatusCreated, bookDto)
}
