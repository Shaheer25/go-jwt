package main

import (
	"net/http"
	"os"

	"github.com/Shaheer25/go-tokenc/router"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3766"
	}

	r := gin.New()
	r.Use(gin.Logger()) // where r=router;

	router.authRouter(r)
	router.userRouter(r)

	r.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Access Granted for Api-1",
		})
	})

	r.GET("/api-2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Access Granted for Api -2",
		})
	})

	r.Run(":" + port)
}
