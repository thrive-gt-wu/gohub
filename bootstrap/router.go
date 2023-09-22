package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	RegisterGlobalMiddleware(router)

	// 注册 api 路由
	routes.RegisterAPIRoutes(router)
}


func RegisterGlobalMiddleware(router *gin.Engine){
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func Setup404Handler(router *gin.Engine){
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 html
			c.String(http.StatusNotFound, "返回页面 404")
		} else {
			// 返回默认 JSON 
			c.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_message": "路由未定义，请确认 URL 和请求方法是否正确",
			})
		}
	
	})
}

