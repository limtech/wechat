package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	TICKET_API_URL string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

type TicketData struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

// get ticket by access_token
func GetTicket(accessToken string) (TicketData, error) {
	var data TicketData

	// get remote data
	res, err := HttpGet(fmt.Sprintf(TICKET_API_URL, accessToken))
	if err != nil {
		return data, err
	}

	// parse json data
	if err := json.Unmarshal(res, &data); err != nil {
		return data, err
	}
	// on success

	if data.ErrCode != 0 {
		return data, errors.New(data.ErrMsg)
	}

	return data, nil
}
