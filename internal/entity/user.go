package entity

import "github.com/gorilla/websocket"

type GetUserInfoReq struct {
	UserId int `json:"userId" binding:"required"`
}

type GetUserInfoResp struct {
	UserId   int    `json:"userId" binding:"required"`
	UserName string `json:"userName" binding:"min=3,max=64"`
	Gender   int    `json:"gender"`
	Age      int    `json:"age" binding:"min=1,max=100"`
}

type (
	LoginReq struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}

	LoginResp struct {
		MallUser
	}
)

// User 用户信息
type (
	User struct {
		ID       string
		Conn     *websocket.Conn
		Send     chan []byte
		Backpack []string
	}
	// Health 用户健康信息
	Health struct {
		ID       string
		IsHealth bool
	}
	// Message 通用信息
	Message struct {
		Action string `json:"action"`
		Data   string `json:"data"`
	}
)
