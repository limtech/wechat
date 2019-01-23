package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// template
	TEMPLATE_SET_INDUSTRY_API string = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=%s"         // 设置所属行业
	TEMPLATE_GET_INDUSTRY_API string = "https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=%s"             // 获取设置的行业信息
	TEMPLATE_GET_ID_API       string = "https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=%s"         // 获得模板ID
	TEMPLATE_GET_ALL_API      string = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s" // 获取模板列表
	TEMPLATE_DELETE_API       string = "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=%s"     // 删除模板
	TEMPLATE_SEND_API         string = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"             // 发送模板消息
)

type (
	Message struct {
		AccessToken string
	}

	MessageTemplateIndustry struct {
		PrimaryIndustry   MessageTemplateIndustryItem `json:"primary_industry"`
		SecondaryIndustry MessageTemplateIndustryItem `json:"secondary_industry"`

		ErrStruct
	}

	MessageTemplateIndustryItem struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
	}

	MessageTemplate struct {
		TemplateId      string `json:"template_id"`
		Title           string `json:"title"`
		PrimaryIndustry string `json:"primary_industry"`
		DeputyIndustry  string `json:"deputy_industry"`
		Content         string `json:"content"`
		Example         string `json:"example"`
	}

	MessageTemplateAll struct {
		TemplateList []MessageTemplate `json:"template_list"`

		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	MessageTemplateSendResult struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
		MsgId   int64  `json:"msgid"`
	}
)

// new Message
func NewMessage(accessToken string) *Message {
	return &Message{
		AccessToken: accessToken,
	}
}

// 设置所属行业
func (self *Message) SetIndustry(industryId1 string, industryId2 string) error {
	res, err := HttpPostJson(
		fmt.Sprintf(TEMPLATE_SET_INDUSTRY_API, self.AccessToken),
		map[string]string{
			"industry_id1": industryId1,
			"industry_id2": industryId2,
		},
		nil,
	)
	if err != nil {
		return err
	}

	var rs ErrStruct
	// parse json data
	if err := json.Unmarshal(res, &rs); err != nil {
		return err
	}

	if rs.ErrCode != 0 {
		return errors.New(rs.ErrMsg)
	}

	return nil
}

// 获取设置的行业信息
func (self *Message) GetIndustry() (MessageTemplateIndustry, error) {
	var data MessageTemplateIndustry
	// get remote data
	res, err := HttpGet(fmt.Sprintf(TEMPLATE_GET_INDUSTRY_API, self.AccessToken))
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

// 获得模板ID
func (self *Message) GetTemplate(templateShortId string) (MessageTemplateIndustry, error) {
	var data MessageTemplateIndustry
	// get remote data
	res, err := HttpGet(fmt.Sprintf(TEMPLATE_GET_INDUSTRY_API, self.AccessToken))
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

// 获取模板列表
func (self *Message) GetTemplateAll() (MessageTemplateAll, error) {
	var data MessageTemplateAll
	// get remote data
	res, err := HttpGet(fmt.Sprintf(TEMPLATE_GET_ALL_API, self.AccessToken))
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

// 删除模板
func (self *Message) DeleteTemplate(templateId string) error {
	res, err := HttpPostJson(
		fmt.Sprintf(TEMPLATE_SEND_API, self.AccessToken),
		map[string]string{
			"template_id": templateId,
		},
		nil,
	)
	if err != nil {
		return err
	}

	rs := ErrStruct{}
	// parse json data
	if err := json.Unmarshal(res, &rs); err != nil {
		return err
	}

	if rs.ErrCode != 0 {
		return errors.New(rs.ErrMsg)
	}

	return nil
}

// 发送模板消息
func (self *Message) SendTemplate(data interface{}) (MessageTemplateSendResult, error) {
	var rs MessageTemplateSendResult

	// get remote data
	res, err := HttpPostJson(fmt.Sprintf(TEMPLATE_SEND_API, self.AccessToken), data, nil)
	if err != nil {
		return rs, err
	}

	// parse json data
	if err := json.Unmarshal(res, &rs); err != nil {
		return rs, err
	}

	if rs.ErrCode != 0 {
		return rs, errors.New(rs.ErrMsg)
	}

	return rs, nil
}
