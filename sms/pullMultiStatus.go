package sms

import (
	"encoding/json"
	"fmt"
	"github.com/sunnyos/tencentSms/context"
	"github.com/sunnyos/tencentSms/utils"
	"time"
)

const PULL_MULTI_STATUS_URL = "https://yun.tim.qq.com/v5/tlssmssvr/pullstatus"

type PullMultiStatus struct {
	Context *context.Context
}

type PullMultiStatusParams struct {
	Max  int    `json:"max"`
	Sig  string `json:"sig"`
	Time int64  `json:"time"`
	Type int    `json:"type"`
}

type PullMultiStatusReply struct {
	Result int                        `json:"result"`
	Errmsg string                     `json:"errmsg"`
	Count  int                        `json:"count"`
	Data   []PullMultiStatusReplyData `json:"data"`
}

type PullMultiStatusReplyData struct {
	ReportStatus    string `json:"report_status"`
	UserReceiveTime string `json:"user_receive_time"`
	Nationcode      string `json:"nationcode"`
	Mobile          string `json:"mobile"`
	Sid             string `json:"sid"`
	Errmsg          string `json:"errmsg"`
	Description     string `json:"description"`
	PullType        int    `json:"pull_type"`
}

func NewPullMultiStatus(context *context.Context) *PullMultiStatus {
	pullMultiStatus := &PullMultiStatus{Context: context}
	return pullMultiStatus
}

func (pull *PullMultiStatus) Fetch(p *PullMultiStatusParams) (*PullMultiStatusReply, error) {
	random := utils.CreateRandom()
	p.Time = time.Now().Unix()
	p.Sig = utils.GetSignatureWithOutMobile(pull.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", PULL_MULTI_STATUS_URL, pull.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	pullMultiStatusReply := &PullMultiStatusReply{}
	err = json.Unmarshal(res, pullMultiStatusReply)
	return pullMultiStatusReply, err
}
