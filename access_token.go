package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ACCESS_TOKEN_API_URL string = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

type AccessTokenData struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

// appid appsecret
func GetAccessToken(appId string, appSecret string) (AccessTokenData, error) {
	var data AccessTokenData
	// get remote data
	res, err := HttpGet(fmt.Sprintf(ACCESS_TOKEN_API_URL, appId, appSecret))
	if err != nil {
		return data, err
	}

	// parse json data
	if err := json.Unmarshal(res, &data); err != nil {
		return data, err
	}

	if data.ErrCode != 0 {
		return data, errors.New(data.ErrMsg)
	}

	return data, nil
}
