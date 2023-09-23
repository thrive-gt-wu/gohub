package routes

import (
	"net/http"
	"gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context)  {
			c.JSON(http.StatusOK, gin.H{
				"hello": "hello",
			})
		})
		authGroup := v1.Group("auth") 
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	
	}
}