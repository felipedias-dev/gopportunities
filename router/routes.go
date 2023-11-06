package router

import (
	"github.com/felipedias-dev/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	v1 := router.Group("api/v1")
	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)

		v1.DELETE("/opening/:id", handler.ExcludeOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}
}
