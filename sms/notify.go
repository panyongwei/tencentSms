package sms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SmsNotify struct {
	UserReceiveTime string `json:"user_receive_time"`
	Nationcode      string `json:"nationcode"`
	Mobile          string `json:"mobile"`
	ReportStatus    string `json:"report_status"`
	Errmsg          string `json:"errmsg"`
	Description     string `json:"description"`
	Sid             string `json:"sid"`
}

type SmsReply struct {
	Time       int64  `json:"time"`
	Nationcode string `json:"nationcode"`
	Mobile     string `json:"mobile"`
	Text       string `json:"text"`
	Sign       string `json:"sign"`
}

// 短信状态回调
func Notify(r *http.Request) ([]SmsNotify, error) {
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	var s []SmsNotify
	e := json.Unmarshal(bodyData, &s)
	return s, e
}

// 短信回复回调
func Reply(r *http.Request) (*SmsReply, error) {
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	reply := &SmsReply{}
	e := json.Unmarshal(bodyData, reply)
	if e != nil {
		return nil, e
	}
	return reply, nil
}
