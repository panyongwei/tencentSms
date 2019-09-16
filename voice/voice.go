package voice

import (
	"github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/context"
)

type TencentVoice struct {
	Context *context.Context
}

type VoiceTel struct {
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type VoiceResult struct {
	Result int    `json:"result"`
	Errmsg string `json:"errmsg"`
	Callid string `json:"callid"`
	Ext    string `json:"ext"`
}

func NewVoice(cfg *config.Config) *TencentVoice {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &TencentVoice{context}
}

func copyConfigToContext(cfg *config.Config, context *context.Context) {
	context.AppId = cfg.AppId
	context.AppKey = cfg.AppKey
	context.Sign = cfg.Sign
}

// 发送语音验证码
func (tv *TencentVoice) NewSendcVoice() *SendcVoice {
	return NewSendcVoice(tv.Context)
}

// 指定模板发送语音通知
func (tv *TencentVoice) NewSendtVoice() *SendtVoice {
	return NewSendtVoice(tv.Context)
}
