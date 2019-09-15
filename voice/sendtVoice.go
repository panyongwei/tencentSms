package voice

import (
	"encoding/json"
	"fmt"
	"github.com/sunnyos/tencentSms/context"
	"github.com/sunnyos/tencentSms/utils"
	"time"
)

const SENDT_VOICE_URL = "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendtvoice"

type SendtVoice struct {
	Context *context.Context
}

type SendtVoiceParams struct {
	TplId     int      `json:"tpl_id"`
	Params    []string `json:"params"`
	PlayTimes int64    `json:"play_times"`
	Sig       string   `json:"sig"`
	Tel       VoiceTel `json:"tel"`
	Time      int64    `json:"time"`
	Ext       string   `json:"ext"`
}

func NewSendtVoice(context *context.Context) *SendtVoice {
	return &SendtVoice{Context: context}
}

func (sv *SendtVoice) Fetch(p *SendtVoiceParams) (*VoiceResult, error) {
	p.Time = time.Now().Unix()
	result, err := sv.send(p)
	return result, err
}

func (sv *SendtVoice) Fetchs(mobile string, nationcode string, tpl_id int, params []string) (*VoiceResult, error) {
	tel := VoiceTel{Mobile: mobile, Nationcode: nationcode}
	p := &SendtVoiceParams{
		TplId:  tpl_id,
		Params: params,
		Time:   time.Now().Unix(),
		Tel:    tel,
	}
	result, err := sv.send(p)
	return result, err
}

func (sv *SendtVoice) send(p *SendtVoiceParams) (*VoiceResult, error) {
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
