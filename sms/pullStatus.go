package sms

import (
	"encoding/json"
	"fmt"
	"github.com/sunnyos/tencentSms/context"
	"github.com/sunnyos/tencentSms/utils"
	"time"
)

const PULL_STATUS_URL = "https://yun.tim.qq.com/v5/tlssmssvr/pullstatus4mobile"

type PullStatus struct {
	Context *context.Context
}

type PullStatusParams struct {
	BeginTime  int64  `json:"begin_time"`
	EndTime    int64  `json:"end_time"`
	Max        int64  `json:"max"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Sig        string `json:"sig"`
	Time       int64  `json:"time"`
	Type       int    `json:"type"`
}

type PullStatusReply struct {
	Result int                   `json:"result"`
	Errmsg string                `json:"errmsg"`
	Count  int                   `json:"count"`
	Data   []PullStatusReplyData `json:"data"`
}

type PullStatusReplyData struct {
	ReportStatus    string `json:"report_status"`
	UserReceiveTime string `json:"user_receive_time"`
	Nationcode      string `json:"nationcode"`
	Mobile          string `json:"mobile"`
	Sid             string `json:"sid"`
	Errmsg          string `json:"errmsg"`
	Description     string `json:"description"`
	PullType        int    `json:"pull_type"`
}

func NewPullStatus(context *context.Context) *PullStatus {
	pullStatus := &PullStatus{Context: context}
	return pullStatus
}

func (pull *PullStatus) Fetch(p *PullStatusParams) (*PullStatusReply, error) {
	random := utils.CreateRandom()
	p.Time = time.Now().Unix()
	p.Sig = utils.GetSignatureWithOutMobile(pull.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", PULL_STATUS_URL, pull.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	pullStatusReply := &PullStatusReply{}
	err = json.Unmarshal(res, pullStatusReply)
	return pullStatusReply, err
}
