package main

import (
	"github.com/emikohmann/ucc-arqsoft-2/ej-memcached-mongo/router"
	"github.com/emikohmann/ucc-arqsoft-2/ej-memcached-mongo/utils/cache"
	"github.com/emikohmann/ucc-arqsoft-2/ej-memcached-mongo/utils/db"
	"github.com/gin-gonic/gin"

	"fmt"
)

var (
	gin_router *gin.Engine
)

func main() {
	gin_router = gin.Default()
	router.MapUrls(gin_router)
	cache.Init_cache()
	err := db.Init_db()
	defer db.Disconect_db()

	if err != nil {
		fmt.Println("Cannot init db")
		fmt.Println(err)
		return
	}
	fmt.Println("Starting server")
	gin_router.Run(":8090")
}
