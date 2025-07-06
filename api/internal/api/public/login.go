package public

type LoginReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

type LoginResp struct {
	Token string `json:"token"`
}
