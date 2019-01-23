package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	TICKET_API string = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

type (
	Ticket struct {
		AccessToken string
	}

	TicketData struct {
		Ticket    string `json:"ticket"`
		ExpiresIn int    `json:"expires_in"`

		ErrStruct
	}
)

// new ticket
func NewTicket(accessToken string) *Ticket {
	return &Ticket{
		AccessToken: accessToken,
	}
}

// get ticket by access_token
func (self *Ticket) GetTicket() (TicketData, error) {
	var data TicketData

	// get remote data
	res, err := HttpGet(fmt.Sprintf(TICKET_API, self.AccessToken))
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
