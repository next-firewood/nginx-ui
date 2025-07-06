package public

type InitServerReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	NginxUrl string `json:"nginxUrl"` // nginx目录
	Node     string `json:"node"`     // 公网IP
}
