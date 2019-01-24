package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// 黑名单管理
	USER_BLACKLIST_ALL_API          string = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s"     // 获取公众号的黑名单列表
	USER_BLACKLIST_BATCH_ADD_API    string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=%s"   // 拉黑用户
	USER_BLACKLIST_BATCH_REMOVE_API string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=%s" // 取消拉黑用户

)

// 获取公众号的黑名单列表，单次最大值为10000
func (self *User) GetBlacklist(nextOpenid string) (UserList, error) {
	var data UserList

	res, err := HttpPostJson(
		fmt.Sprintf(USER_BLACKLIST_ALL_API, self.AccessToken),
		map[string]string{
			"begin_openid": nextOpenid,
		},
		nil,
	)
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

// 获取公众号的黑名单列表(所有)
func (self *User) GetBlacklistAll() (*UserListAll, error) {
	list := &UserListAll{}
	nextOpenid := ""

	for {
		data, err := self.GetBlacklist(nextOpenid)
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

// 批量黑名单操作
func (self *User) BlacklistBatchAction(openid []string, action string) error {
	data := map[string]interface{}{
		"openid_list": openid,
	}

	url := ""
	switch action {
	case "add":
		url = USER_BLACKLIST_BATCH_ADD_API
	case "remove":
		url = USER_BLACKLIST_BATCH_REMOVE_API
	default:
		url = USER_BLACKLIST_BATCH_ADD_API
	}
	// get remote data
	res, err := HttpPostJson(fmt.Sprintf(url, self.AccessToken), data, nil)
	if err != nil {
		return err
	}

	rs := ErrStruct{}
	if err := json.Unmarshal(res, &rs); err != nil {
		return err
	}

	if rs.ErrCode != 0 {
		return errors.New(rs.ErrMsg)
	}

	return nil
}

// 批量拉黑用户
func (self *User) BlacklistBatchAdd(openid []string) error {
	return self.BlacklistBatchAction(openid, "add")
}

// 批量拉黑用户
func (self *User) BlacklistBatchRemove(openid []string) error {
	return self.BlacklistBatchAction(openid, "remove")
}
