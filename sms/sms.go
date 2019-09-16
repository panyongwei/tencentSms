package sms

import (
	"github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/context"
)

type TencentSms struct {
	Context *context.Context
}

type SmsTel struct {
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

func NewSms(cfg *config.Config) *TencentSms {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &TencentSms{context}
}

func copyConfigToContext(cfg *config.Config, context *context.Context) {
	context.AppId = cfg.AppId
	context.AppKey = cfg.AppKey
	context.Sign = cfg.Sign
}

// 指定模板单发短信
func (t *TencentSms) GetSmsSender() *SendSms {
	return NewSendSms(t.Context)
}

// 指定模板群发短信
func (t *TencentSms) GetSendMultiSms() *SendMultiSms {
	return NewMultiSms(t.Context)
}

// 拉取短信状态
func (t *TencentSms) GetPullMultiStatus() *PullMultiStatus {
	return NewPullMultiStatus(t.Context)
}

// 拉取单个手机短信状态
func (t *TencentSms) GetPullStatus() *PullStatus {
	return NewPullStatus(t.Context)
}
