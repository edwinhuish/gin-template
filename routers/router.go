package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	apis "gin-template/api/v1"
	"gin-template/middleware"

	"github.com/gin-contrib/pprof"
)

// 初始化总路由

func InitRouters() *gin.Engine {
	r := gin.Default()

	pprof.Register(r)

	r.Use(middleware.CorsByRules())
	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	r.StaticFS("/cache", http.Dir("cache"))

	apiv1Group(r)

	return r
}

func apiv1Group(r *gin.Engine) {

	apiv1 := r.Group("/api/v1", middleware.Timeout())
	{

		exampleApi := apis.NewExampleController()

		apiv1.GET("/example/hello", exampleApi.Hello)
		apiv1.POST("/example/upload", exampleApi.Upload)
	}
}
