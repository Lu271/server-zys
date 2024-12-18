package router

import (
	"github.com/Lu271/server-zys/api/controller"
	"github.com/Lu271/server-zys/api/middleware"
	"github.com/Lu271/server-zys/httputils"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
)

func RegisterRouter(router *gin.Engine) {
	router.Use(
		middleware.Qps,
		middleware.Delay,
	)
	router.GET("/metrics", controller.Metrics)

	// ws
	router.GET("/ws", controller.InitHandler)

	router.Use(
		middleware.RecoverMiddleware,
		middleware.CheckLogin,
		middleware.AccessLogger,
		middleware.Context,
		middleware.SentinelMiddleware(),
	)

	// 管理后台相关路由
	admin := router.Group("/admin")
	RegisterAdminRouter(admin)
	// 用户侧路由
	api := router.Group("/api")
	RegisterApiRouter(api)

}

func RegisterAdminRouter(router *gin.RouterGroup) {
	router.PUT("/users", controller.GetUserInfo)
}

func RegisterApiRouter(router *gin.RouterGroup) {
	router.GET("/healthCheck", func(c *gin.Context) {
		r := rand.Intn(10000)
		c.JSON(http.StatusOK, httputils.SuccessWithData(r))
		return
	})
	router.POST("/login", controller.UserLogin)
	router.GET("/say/hello", controller.SayHello)
}
