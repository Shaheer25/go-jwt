package router

import "github.com/gin-gonic/gin"

func authRouter(incomingRoutes *gin.Engine){
	incomingRoutes.POST("/user/signup",controller.signup())
	incomingRoutes.POST("/user/login",controller.login())
}