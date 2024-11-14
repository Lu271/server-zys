package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-zys/httputils"
	"server-zys/internal/entity"
	"server-zys/internal/service"
)

func GetUserInfo(c *gin.Context) {
	var req = entity.GetUserInfoReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, httputils.Error(err))
		return
	}
	c.JSON(http.StatusOK, httputils.SuccessWithData(service.GetUserInfo(req.UserId)))
}
