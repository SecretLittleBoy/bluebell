package models

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamRefreshToken struct {
	AcessToken  string `json:"access_token" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"` // 帖子id
	Direction int8   `json:"direction" binding:"oneof=1 0 -1"` // 赞成票(1)还是反对票(-1)还是取消投票(0)
}