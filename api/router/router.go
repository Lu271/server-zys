package router

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"server-zys/api/controller"
	"server-zys/api/middleware"
	"server-zys/httputils"
)

func RegisterRouter(router *gin.Engine) {
	router.Use(
		middleware.Qps,
		middleware.Delay,
	)
	router.GET("/metrics", controller.Metrics)

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
}
