package server

import (
	"github.com/gin-gonic/gin"
	"journey/server/handlers"
	"journey/server/middlewares"
)

func NewServer() *gin.Engine {
	r := gin.Default()

	protected := r.Group("/")
	protected.Use(middleware.FakeAuthMiddleware())

	protected.GET("/accepted", handlers.GetAcceptedJourneysHandler)
	protected.GET("/created", handlers.GetCreatedJourneysHandler)
	protected.POST("/createJourney", handlers.CreateJourneyHandler)

	return r
}
