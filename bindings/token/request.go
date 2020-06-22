package token

type RequestReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestRsp struct {
	Token    string `json:"token"`
	ExpireAt int    `json:"expire_at"`
}
