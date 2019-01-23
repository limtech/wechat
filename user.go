package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// user
	USER_LIST_API       string = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s"        //
	USER_INFO_API       string = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN" // 获取用户基本信息（包括UnionID机制）
	USER_INFO_BATCH_API string = "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=%s"             // 批量获取用户基本信息
	USRE_UPDATE_REMARK  string = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=%s"         // 设置用户备注名

	// user.tag
	USER_TAG_LIST_API            string = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"                    // 获取公众号已创建的标签
	USER_TAG_CREATE_API          string = "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"                 // 创建标签
	USER_TAG_UPDATE_API          string = "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s"                 // 编辑标签
	USER_TAG_DELETE_API          string = "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s"                 // 删除标签
	USER_TAG_USERS_API           string = "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s"                // 获取标签下粉丝列表
	USER_TAG_BATCH_TAGGING_API   string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s"   // 批量为用户打标签
	USER_TAG_BATCH_UNTAGGING_API string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s" // 批量为用户取消标签
	USER_TAG_ID_LIST_API         string = "https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=%s"              // 获取用户身上的标签列表

	// 黑名单管理
	USER_BLACKLIST_ALL_API          string = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s"     // 获取公众号的黑名单列表
	USER_BLACKLIST_BATCH_ADD_API    string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=%s"   // 拉黑用户
	USER_BLACKLIST_BATCH_REMOVE_API string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=%s" // 取消拉黑用户

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

	UserTagItem struct {
		Id    int64  `json:"id"`
		Name  string `json:"name"`
		Count int64  `json:"count"`
	}

	UserTagData struct {
		Tag UserTagItem `json:"tag"`

		ErrStruct
	}

	UserTagList struct {
		Tags []UserTagItem `json:"tags"`

		ErrStruct
	}

	UserTagUserListData struct {
		Count int64 `json:"count"`
		Data  struct {
			Openid []string `json:"openid"`
		} `json:"data"`
		NextOpenid string `json:"next_openid"`

		ErrStruct
	}

	UserTagIdList struct {
		TagidList []int64 `json:"tagid_list"`

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

// 创建标签
func (self *User) CreateTag(name string) (UserTagData, error) {
	data := UserTagData{
		Tag: UserTagItem{
			Name: name,
		},
	}
	// get remote data
	res, err := HttpPostJson(fmt.Sprintf(USER_TAG_CREATE_API, self.AccessToken), data, nil)
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

// 获取公众号已创建的标签
func (self *User) GetTagList() (UserTagList, error) {
	var data UserTagList

	// get remote data
	res, err := HttpGet(fmt.Sprintf(USER_TAG_LIST_API, self.AccessToken))
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

// 编辑标签
func (self *User) UpdateTag(id int64, name string) (UserTagData, error) {
	data := UserTagData{
		Tag: UserTagItem{
			Id:   id,
			Name: name,
		},
	}
	// get remote data
	res, err := HttpPostJson(fmt.Sprintf(USER_TAG_UPDATE_API, self.AccessToken), data, nil)
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

// 删除标签
func (self *User) DeleteTag(id int64) error {
	data := UserTagData{
		Tag: UserTagItem{
			Id: id,
		},
	}
	// get remote data
	res, err := HttpPostJson(fmt.Sprintf(USER_TAG_DELETE_API, self.AccessToken), data, nil)
	if err != nil {
		return err
	}

	// parse json data
	if err := json.Unmarshal(res, &data); err != nil {
		return err
	}

	if data.ErrCode != 0 {
		return errors.New(data.ErrMsg)
	}

	return nil
}

// 获取标签下粉丝列表
func (self *User) GetTagUsers(tagId int64, nextOpenid string) (UserTagUserListData, error) {
	data := UserTagUserListData{}

	// get remote data
	res, err := HttpPostJson(
		fmt.Sprintf(USER_TAG_USERS_API, self.AccessToken),
		map[string]interface{}{
			"tagid":       tagId,
			"next_openid": nextOpenid,
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

// 批量为用户打标签
func (self *User) BatchTagging(tagId int64, openid []string) error {
	// get remote data
	res, err := HttpPostJson(
		fmt.Sprintf(USER_TAG_BATCH_TAGGING_API, self.AccessToken),
		map[string]interface{}{
			"tagid":       tagId,
			"openid_list": openid,
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

// 批量为用户取消标签
func (self *User) BatchUnTagging(tagId int64, openid []string) error {
	// get remote data
	res, err := HttpPostJson(
		fmt.Sprintf(USER_TAG_BATCH_UNTAGGING_API, self.AccessToken),
		map[string]interface{}{
			"tagid":       tagId,
			"openid_list": openid,
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

// 获取用户身上的标签列表
func (self *User) GetUserTagIds(openid string) (UserTagIdList, error) {
	var data UserTagIdList

	res, err := HttpPostJson(
		fmt.Sprintf(USER_TAG_ID_LIST_API, self.AccessToken),
		map[string]string{
			"openid": openid,
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
