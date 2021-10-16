package main

import (
	"demo/blog/config"
	"demo/blog/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	routers.InitRoutes(r)

	listenPort := config.GetConfig().GetString("listen")
	http.ListenAndServe(listenPort, r)
}
