package main

import (
	_ "fmt"

	"net/http"

	"github.com/gin-gonic/gin"

	// _ swaggerFiles "github.com/swaggo/files"
	// _ ginSwagger "github.com/swaggo/gin-swagger"
	control "gin/Control"
	"gin/model"
)

func main() {
	router := gin.Default()
	member := control.NewMember()

	// router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/getAllData", func(c *gin.Context) {
		c.JSON(http.StatusOK, member.GetAllData())
	})
	router.POST("/postCreateData", func(c *gin.Context) {
		var v model.Member
		err := c.ShouldBind(&v)
		if err != nil {
			c.JSON(http.StatusExpectationFailed, gin.H{
				"message": "error format",
			})
			return
		}
		member.PostCreateData(v)
		c.JSON(http.StatusOK, v)
	})
	router.Run("localhost:3000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
