package delivery

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(clubHandler *ClubHandler) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	clubHandler.RegisterRoutes(api)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
