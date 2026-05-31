package router

import (
	"github.com/gin-gonic/gin"

	"github.com/georgigeorgiev/cloudops/auth-service/internal/handler"
	"github.com/georgigeorgiev/cloudops/auth-service/internal/middleware"
)

func Setup(
	authHandler *handler.AuthHandler,
	jwtSecret string,
) *gin.Engine {

	router := gin.Default()

	router.GET(
		"/health",
		func(c *gin.Context) {

			c.JSON(
				200,
				gin.H{
					"status": "UP",
				},
			)
		},
	)

	auth := router.Group("/auth")

	{
		auth.POST(
			"/register",
			authHandler.Register,
		)

		auth.POST(
			"/login",
			authHandler.Login,
		)
	}

	protected := router.Group("/")

	protected.Use(
		middleware.JWTAuth(jwtSecret),
	)

	protected.GET(
		"/me",
		authHandler.Me,
	)

	return router
}
