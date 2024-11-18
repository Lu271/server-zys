package entity

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
