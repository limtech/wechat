package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// user.tag
	USER_TAG_LIST_API            string = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"                    // 获取公众号已创建的标签
	USER_TAG_CREATE_API          string = "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"                 // 创建标签
	USER_TAG_UPDATE_API          string = "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s"                 // 编辑标签
	USER_TAG_DELETE_API          string = "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s"                 // 删除标签
	USER_TAG_USERS_API           string = "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s"                // 获取标签下粉丝列表
	USER_TAG_BATCH_TAGGING_API   string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s"   // 批量为用户打标签
	USER_TAG_BATCH_UNTAGGING_API string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s" // 批量为用户取消标签
	USER_TAG_ID_LIST_API         string = "https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=%s"              // 获取用户身上的标签列表
)

type (
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
