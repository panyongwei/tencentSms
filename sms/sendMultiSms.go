package sms

import (
	"encoding/json"
	"fmt"
	"github.com/sunnyos/tencentSms/context"
	"github.com/sunnyos/tencentSms/utils"
	"strings"
	"time"
)

const MULTI_SMS_URL = "https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2"

type SendMultiSms struct {
	Context *context.Context
}

type MultiParams struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    []SmsTel `json:"tel"`
	Time   int64    `json:"time"`
	TplId  int64    `json:"tpl_id"`
}

type MultiResult struct {
	Result int                 `json:"result"`
	Errmsg string              `json:"errmsg"`
	Ext    string              `json:"ext"`
	Detail []MultiResultDetail `json:"detail"`
}

type MultiResultDetail struct {
	Errmsg     string `json:"errmsg"`
	Fee        int    `json:"fee"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
	Result     int    `json:"result"`
	Sid        string `json:"sid"`
}

func NewMultiSms(context *context.Context) *SendMultiSms {
	sendMultiSms := &SendMultiSms{Context: context}
	return sendMultiSms
}

func (sms *SendMultiSms) Fetch(p *MultiParams) (*MultiResult, error) {
	p.Time = time.Now().Unix()
	p.Sign = sms.Context.Sign
	var mobile string
	for _, v := range p.Tel {
		mobile += v.Mobile + ","
	}
	mobile = strings.TrimRight(mobile, ",")
	result, err := sms.send(p, mobile)
	return result, err
}

func (sms *SendMultiSms) Fetchs(mobile []string, nationcode []string, params []string, tplId int64) (*MultiResult, error) {
	p := &MultiParams{
		Time:  time.Now().Unix(),
		Sign:  sms.Context.Sign,
		TplId: tplId,
	}
	t := make([]SmsTel, 0)
	var m string
	for i, v := range mobile {
		m += v + ","
		t = append(t, SmsTel{Mobile: v, Nationcode: nationcode[i]})
	}
	p.Tel = t
	m = strings.TrimRight(m, ",")
	result, err := sms.send(p, m)
	return result, err
}

func (sms *SendMultiSms) send(p *MultiParams, mobile string) (*MultiResult, error) {
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithMobile(sms.Context.AppKey, mobile, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", MULTI_SMS_URL, sms.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &MultiResult{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return nil, err
	}
	result.Errmsg = utils.CodeMsg[result.Result]
	return result, nil
}
