package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	USER_LIST_API_URL string = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s"
	USER_INFO_API_URL string = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
)

type UserList struct {
	Total  int64
	Openid []string
}

type UserListData struct {
	Total int64 `json:"total"`
	Count int64 `json:"count"`
	Data  struct {
		Openid []string `json:"openid"`
	} `json:"data"`
	NextOpenid string `json:"next_openid"`

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type UserInfo struct {
	Subscribe      int     `json:"subscribe"`
	Openid         string  `json:"openid"`
	Nickname       string  `json:"nickname"`
	Sex            int     `json:"sex"`
	Language       string  `json:"language"`
	City           string  `json:"city"`
	Provice        string  `json:"province"`
	Country        string  `json:"country"`
	HeadImgUrl     string  `json:"headimgurl"`
	SubscribeTime  int64   `json:"subscribe_time"`
	UnionId        string  `json:"unionid"`
	Remark         string  `json:"remark"`
	GroupId        int64   `json:"groupid"`
	TagidList      []int64 `json:"tagid_list"`
	SubscribeScene string  `json:"subscribe_scene"`
	QrScene        int64   `json:"qr_scene"`
	QrSceneStr     string  `json:"qr_scene_str"`

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func GetUserList(accessToken string, nextOpenid string) (UserListData, error) {
	var data UserListData
	// get remote data
	res, err := HttpGet(fmt.Sprintf(USER_LIST_API_URL, accessToken, nextOpenid))
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

func GetUserListAll(accessToken string) (*UserList, error) {
	list := &UserList{}
	nextOpenid := ""

	for {
		data, err := GetUserList(accessToken, nextOpenid)
		if err != nil {
			continue
		}

		nextOpenid = data.NextOpenid
		// end for
		if data.Count == 0 {
			break
		}

		list.Openid = append(list.Openid, data.Data.Openid...)
	}

	return list, nil
}

func GetUserInfo(accessToken string, openid string) (UserInfo, error) {
	var data UserInfo

	// get remote data
	res, err := HttpGet(fmt.Sprintf(USER_INFO_API_URL, accessToken, openid))
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
