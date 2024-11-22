package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"server-zys/httputils"
	"server-zys/internal/entity"
	"server-zys/internal/service"
	"server-zys/logs"
)

func SayHello(c *gin.Context) {
	resp, err := service.SayHello(c)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func GetUserInfo(c *gin.Context) {
	var req = entity.GetUserInfoReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logs.Error(c, fmt.Sprintf("get user info err: %v", err.Error()))
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	logs.Info(c, "get user info success")
	c.JSON(http.StatusOK, httputils.SuccessWithData(service.GetUserInfo(req.UserId)))
}

func UserLogin(c *gin.Context) {
	var req = entity.LoginReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logs.Error(c, fmt.Sprintf("get login info err: %v", err.Error()))
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}

	resp, err := service.UserLogin(c, req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(resp))
}

func Metrics(c *gin.Context) {
	handler := promhttp.Handler()
	handler.ServeHTTP(c.Writer, c.Request)
}
