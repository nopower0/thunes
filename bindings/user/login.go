package user

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRsp struct {
	UID      int    `json:"uid"`
	Username string `json:"username"`
}
