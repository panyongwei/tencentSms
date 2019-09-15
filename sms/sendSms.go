package sms

import (
	"encoding/json"
	"fmt"
	"github.com/sunnyos/tencentSms/context"
	"github.com/sunnyos/tencentSms/utils"
	"time"
)

const SMS_URL = "https://yun.tim.qq.com/v5/tlssmssvr/sendsms"

type SendSms struct {
	Context *context.Context
}

type Params struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    SmsTel   `json:"tel"`
	Time   int64    `json:"time"`
	TplId  int64    `json:"tpl_id"`
}

type Result struct {
	Result int    `json:"result"`
	Errmsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Fee    int    `json:"fee"`
	Sid    string `json:"sid"`
}

func NewSendSms(context *context.Context) *SendSms {
	smsSingleSender := &SendSms{Context: context}
	return smsSingleSender
}

func (sms *SendSms) Fetch(p *Params) (*Result, error) {
	p.Time = time.Now().Unix()
	result, err := sms.send(p)
	return result, err
}

func (sms *SendSms) Fetchs(mobile string, nationcode string, params []string, tplId int64) (*Result, error) {
	p := &Params{
		Tel:    SmsTel{Mobile: mobile, Nationcode: nationcode},
		Time:   time.Now().Unix(),
		Sign:   sms.Context.Sign,
		Params: params,
		TplId:  tplId,
	}
	result, err := sms.send(p)
	return result, err
}

func (sms *SendSms) send(p *Params) (*Result, error) {
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithMobile(sms.Context.AppKey, p.Tel.Mobile, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", SMS_URL, sms.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &Result{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}
	result.Errmsg = utils.CodeMsg[result.Result]
	return result, nil
}
