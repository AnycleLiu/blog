package routers

import (
	"demo/blog/controllers"

	"github.com/gin-gonic/gin"
)

func SetArticleRoutes(r *gin.RouterGroup) {
	g := r.Group("/article")
	{
		a := new(controllers.ArticleController)

		g.POST("", a.Create)
		g.GET("/:id", a.Get)

	}
}
