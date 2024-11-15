package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"server-zys/httputils"
	"server-zys/internal/entity"
	"server-zys/internal/service"
	"server-zys/logs"
)

func GetUserInfo(c *gin.Context) {
	var req = entity.GetUserInfoReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		logs.Error(fmt.Sprintf("get user info err: %v", err.Error()))
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	logs.Info("get user info success")
	c.JSON(http.StatusOK, httputils.SuccessWithData(service.GetUserInfo(req.UserId)))
}
