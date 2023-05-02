package responses

type LoginResponse struct {
	Sessionkey string `json:"sessionkey"`
	Error      *Err   `json:"error"`
	Avatar     string `json:"avatar"`
	Uid        int64  `json:"uid"`
}
