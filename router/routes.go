package router

import (
	docs "github.com/felipedias-dev/gopportunities/docs"
	"github.com/felipedias-dev/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)

		v1.DELETE("/opening/:id", handler.ExcludeOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
