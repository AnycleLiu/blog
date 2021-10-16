package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	SetArticleRoutes(v1)
}
