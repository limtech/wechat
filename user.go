package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// user
	USER_LIST_API       string = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s"        // 获取用户列表
	USER_INFO_API       string = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN" // 获取用户基本信息（包括UnionID机制）
	USER_INFO_BATCH_API string = "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=%s"             // 批量获取用户基本信息
	USRE_UPDATE_REMARK  string = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=%s"         // 设置用户备注名
)

type (
	User struct {
		AccessToken string
	}

	UserList struct {
		Total int64 `json:"total"`
		Count int64 `json:"count"`
		Data  struct {
			Openid []string `json:"openid"`
		} `json:"data"`
		NextOpenid string `json:"next_openid"`

		ErrStruct
	}

	UserListAll struct {
		Total  int64    `json:"total"`
		Openid []string `json:"openid"`
	}

	UserInfo struct {
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
	}

	UserInfoData struct {
		UserInfo
		ErrStruct
	}

	UserInfoBatch struct {
		UserInfoList []UserInfo `json:"user_info_list"`
		ErrStruct
	}
)

// new user
func NewUser(accessToken string) *User {
	return &User{
		AccessToken: accessToken,
	}
}

// 获取用户列表，单次最大值为10000
func (self *User) GetList(nextOpenid string) (UserList, error) {
	var data UserList
	// get remote data
	res, err := HttpGet(fmt.Sprintf(USER_LIST_API, self.AccessToken, nextOpenid))
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

// 获取所有用户列表
func (self *User) GetListAll() (*UserListAll, error) {
	list := &UserListAll{}
	nextOpenid := ""

	for {
		data, err := self.GetList(nextOpenid)
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

// 获取用户基本信息(UnionID机制)
func (self *User) GetUserInfo(openid string) (UserInfoData, error) {
	var data UserInfoData
	res, err := HttpGet(fmt.Sprintf(USER_INFO_API, self.AccessToken, openid))
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

// 批量获取用户基本信息
func (self *User) GetUserInfoBatch(openid []string) (UserInfoBatch, error) {
	var rs UserInfoBatch
	type postItem struct {
		Openid string `json:"openid"`
		Lang   string `json:"lang"`
	}
	type postData struct {
		UserList []postItem `json:"user_list"`
	}

	data := &postData{}
	for _, v := range openid {
		data.UserList = append(data.UserList, postItem{
			Openid: v,
			Lang:   "zh_CN",
		})
	}

	res, err := HttpPostJson(fmt.Sprintf(USRE_UPDATE_REMARK, self.AccessToken), data, nil)
	if err != nil {
		return rs, err
	}

	if err := json.Unmarshal(res, &data); err != nil {
		return rs, err
	}

	if rs.ErrCode != 0 {
		return rs, errors.New(rs.ErrMsg)
	}

	return rs, nil
}

// 设置用户备注名
func (self *User) UpdateRemark(openid string, remark string) error {
	// get remote data
	res, err := HttpPostJson(
		fmt.Sprintf(USRE_UPDATE_REMARK, self.AccessToken),
		map[string]interface{}{
			"tagid":  openid,
			"remark": remark,
		},
		nil,
	)
	if err != nil {
		return err
	}

	data := ErrStruct{}
	// parse json data
	if err := json.Unmarshal(res, &data); err != nil {
		return err
	}

	if data.ErrCode != 0 {
		return errors.New(data.ErrMsg)
	}

	return nil
}
