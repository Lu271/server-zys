package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"io"
	"log"
	"net/http"
	"runtime"
	"server-zys/httputils"
	"time"
)

const size = 64 << 10

func Context(c *gin.Context) {
	traceId := c.GetHeader("traceId")
	if len(traceId) <= 0 {
		traceId = uuid.New().String()
	}
	c.Set("traceId", traceId)
	c.Set("startTime", time.Now().UnixNano())
}

func CheckLogin(c *gin.Context) {
	token := c.GetHeader("mall-auth-token")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, httputils.Error(httputils.UserNotLogin))
		return
	}
	c.Set("userId", token)
}

func RecoverMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 打印堆栈信息
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("http: panic serving err: %v\n%s", r, buf)
			c.JSON(http.StatusOK, httputils.Error(httputils.InterNalError))
			c.Abort()
		}
	}()
	c.Next()
}

func AccessLogger(c *gin.Context) {
	// 开始时间
	startTime := time.Now()
	var body []byte
	if c.Request.Body != nil {
		body, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	c.Next()
	endTime := time.Now()
	latencyTime := endTime.Sub(startTime)
	reqMethod := c.Request.Method
	reqUri := c.Request.RequestURI
	statusCode := c.Writer.Status()

	msg := fmt.Sprintf("method:%v uri:%v req_body:%v status_code:%v latency:%v",
		reqMethod, reqUri, cast.ToString(body), statusCode, latencyTime)
	fmt.Println(msg)
}
