package controller

import (
	"github.com/Lu271/server-zys/internal/socket"
	"github.com/gin-gonic/gin"
)

func InitHandler(c *gin.Context) {
	// 协议升级
	user := socket.Upgrade(c)

	// 开启读取 Goroutine
	go socket.Read(user)

	// 开启发送 Goroutine
	go socket.Write(user)
}
