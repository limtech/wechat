package wechat

import (
	"crypto/sha1"
	"fmt"
	"time"
)

const (
	JSSDK_SIGNATURE_STRING string = "jsapi_ticket=%s&noncestr=%s&timestamp=%d&url=%s"
)

type JSSDK struct {
	AppId     string `json:"appId"`
	Timestamp int64  `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Signature string `json:"signature"`
}

func GetJSSDKSignature(ticket string, url string) (signature string, nonceStr string, timestamp int64) {
	nonceStr = RandomString()
	timestamp = time.Now().Unix()
	signature = fmt.Sprintf("%x", sha1.Sum([]byte(fmt.Sprintf(JSSDK_SIGNATURE_STRING, ticket, nonceStr, timestamp, url))))
	return signature, nonceStr, timestamp
}

func GetJSSDKConfig(appId string, ticket string, url string) JSSDK {
	signature, nonceStr, timestamp := GetJSSDKSignature(ticket, url)
	return JSSDK{
		AppId:     appId,
		Timestamp: timestamp,
		NonceStr:  nonceStr,
		Signature: signature,
	}
}
