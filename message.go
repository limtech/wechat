package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	MESSAGE_TEMPLATE_ALL_API_URL  string = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=%s"
	MESSAGE_TEMPLATE_SEND_API_URL string = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"
)

type MessageTemplate struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

type MessageTemplateAll struct {
	TemplateList []MessageTemplate `json:"template_list"`

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type MessageTemplateSendResult struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgId   int64  `json:"msgid"`
}

func GetMessageTemplateAll(accessToken string) (MessageTemplateAll, error) {
	var data MessageTemplateAll
	// get remote data
	res, err := HttpGet(fmt.Sprintf(MESSAGE_TEMPLATE_ALL_API_URL, accessToken))
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

func SendMessageTemplate(accessToken string, jsonStruct interface{}) (MessageTemplateSendResult, error) {
	var data MessageTemplateSendResult
	// get remote data
	res, err := HttpPostJson(fmt.Sprintf(MESSAGE_TEMPLATE_ALL_API_URL, accessToken), jsonStruct)
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
