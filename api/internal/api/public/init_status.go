package public

type InitStatusRes struct {
	Status int32 `json:"status"` // 1-初始化 2-未初始化账号 3-未初始化nginx目录
}
