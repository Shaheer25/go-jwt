package router

import (
	"github.com/Shaheer25/go-tokenc/middleware"
	"github.com/gin-gonic/gin"
)

func userRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/user", controller.getUser())
	incomingRoutes.GET("/user/:id", controller.getUserById())
}

