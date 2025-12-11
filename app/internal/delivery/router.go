package delivery

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	clubHandler *ClubHandler,
	teamHandler *TeamHandler,
	coachHandler *CoachHandler,
	playerHandler *PlayerHandler,
	stadiumHandler *StadiumHandler,
	staffHandler *StaffHandler,
) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	clubHandler.RegisterRoutes(api)
	teamHandler.RegisterRoutes(api)
	coachHandler.RegisterRoutes(api)
	playerHandler.RegisterRoutes(api)
	stadiumHandler.RegisterRoutes(api)
	staffHandler.RegisterRoutes(api)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
