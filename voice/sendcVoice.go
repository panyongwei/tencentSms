package voice

import (
	"encoding/json"
	"fmt"
	"github.com/sunnyos/tencentSms/context"
	"github.com/sunnyos/tencentSms/utils"
	"time"
)

const SENDC_VOICE_URL = "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendcvoice"

type SendcVoice struct {
	Context *context.Context
}

type SendcVoiceParams struct {
	Ext       string   `json:"ext"`
	Msg       string   `json:"msg"`
	PlayTimes int64    `json:"play_times"`
	Sig       string   `json:"sig"`
	Tel       VoiceTel `json:"tel"`
	Time      int64    `json:"time"`
}

func NewSendcVoice(context *context.Context) *SendcVoice {
	return &SendcVoice{Context: context}
}

func (sv *SendcVoice) Fetch(p *SendcVoiceParams) (*VoiceResult, error) {
	random := utils.CreateRandom()
	p.Time = time.Now().Unix()
	p.Sig = utils.GetSignatureWithMobile(sv.Context.AppKey, p.Tel.Mobile, p.Time, random)
	result, err := sv.send(p)
	return result, err
}

func (sv *SendcVoice) Fetchs(msg string, mobile string, nationcode string) (*VoiceResult, error) {
	tel := VoiceTel{Mobile: mobile, Nationcode: nationcode}
	p := &SendcVoiceParams{
		Msg:  msg,
		Time: time.Now().Unix(),
		Tel:  tel,
	}
	result, err := sv.send(p)
	return result, err
}

func (sv *SendcVoice) send(p *SendcVoiceParams) (*VoiceResult, error) {
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithMobile(sv.Context.AppKey, p.Tel.Mobile, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", SENDC_VOICE_URL, sv.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	voiceResult := &VoiceResult{}
	err = json.Unmarshal(res, voiceResult)
	return voiceResult, err
}
