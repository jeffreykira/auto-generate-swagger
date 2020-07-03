package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jeffreykira/log-management/service/server/api/controller"
	"github.com/jeffreykira/log-management/service/server/api/middleware"

	ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/gin-swagger/swaggerFiles"
	swaggerFiles "github.com/swaggo/files"
)

//NewRouter ...
func NewRouter() *gin.Engine {
	gin.DisableConsoleColor()

	r := gin.Default()
	c := controller.NewController()

	url := ginSwagger.URL("http://localhost:5608/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.POST("/api/v1/auth/signin", c.SigninValid)
	authorized := r.Group("/api/v1")
	authorized.Use(middleware.ValidateToken)
	{
		authorized.GET("/test", controller.TestHandler)
		// authorized.GET("/system/status", c.CheckStatus)
	}

	return r
}
