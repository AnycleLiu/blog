package main

import (
	"demo/blog/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routers.InitRoutes(r)

	http.ListenAndServe(":8080", r)
}
