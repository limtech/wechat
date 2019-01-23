package wechat

type ErrStruct struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
