package voice

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Voicecode struct {
	VoicecodeCallback VoiceCallback `json:"voicecode_callback"`
}

type Voiceprompt struct {
	VoicepromptCallback VoiceCallback `json:"voiceprompt_callback"`
}

type VoiceCallback struct {
	Result        string `json:"result"`
	AcceptTime    string `json:"accept_time"`
	CallFrom      string `json:"call_from"`
	Callid        string `json:"callid"`
	EndCalltime   string `json:"end_calltime"`
	Fee           string `json:"fee"`
	Mobile        string `json:"mobile"`
	Nationcode    string `json:"nationcode"`
	StartCalltime string `json:"start_calltime"`
}

type Voicekey struct {
	VoicekeyCallback VoicekeyCallback `json:"voicekey_callback"`
}

type VoicekeyCallback struct {
	CallFrom   string `json:"call_from"`
	Callid     string `json:"callid"`
	Keypress   string `json:"keypress"`
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type VoiceFailure struct {
	VoiceFailureCallback VoiceFailureCallback `json:"voice_failure_callback"`
}
type VoiceFailureCallback struct {
	CallFrom      string `json:"call_from"`
	Callid        string `json:"callid"`
	FailureCode   int    `json:"failure_code"`
	FailureReason string `json:"failure_reason"`
	Mobile        string `json:"mobile"`
	Nationcode    string `json:"nationcode"`
}

// 语音验证码状态通知
func VoiceCodeNotify(r *http.Request) (*Voicecode, error) {
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	voicecode := &Voicecode{}
	e := json.Unmarshal(bodyData, voicecode)
	if e != nil {
		return nil, e
	}
	return voicecode, nil
}

// 语音通知状态通知
func VoicePromptNotify(r *http.Request) (*Voiceprompt, error) {
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	voiceprompt := &Voiceprompt{}
	e := json.Unmarshal(bodyData, voiceprompt)
	if e != nil {
		return nil, e
	}
	return voiceprompt, nil
}

// 语音通知按键通知
func VoiceKeyNotify(r *http.Request) (*Voicekey, error) {
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	voicekey := &Voicekey{}
	e := json.Unmarshal(bodyData, voicekey)
	if e != nil {
		return nil, e
	}
	return voicekey, nil
}

// 语音送达失败原因推送
func VoiceFailureNotify(r *http.Request) (*VoiceFailure, error) {
	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	voiceFailure := &VoiceFailure{}
	e := json.Unmarshal(bodyData, voiceFailure)
	if e != nil {
		return nil, e
	}
	return voiceFailure, nil
}
