package examples

import (
	"fmt"
	"github.com/sunnyos/tencentSms/voice"
	"net/http"
)

// 发送语音验证码
func TestSendVoice() {
	c := &voice.Config{AppId: "", AppKey: "", Sign: ""}
	v := voice.NewVoice(c)

	// demo1
	tel := voice.VoiceTel{Mobile: "1300000000", Nationcode: "86"}
	p := &voice.SendcVoiceParams{Msg: "1122", Tel: tel}
	result, err := v.NewSendcVoice().Fetch(p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// demo2
	result1, err1 := v.NewSendcVoice().Fetchs("1233", "13000000000", "86")
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}
}

// 指定模板发送语音通知
func TestSendtVoice() {
	c := &voice.Config{AppId: "", AppKey: "", Sign: ""}
	v := voice.NewVoice(c)

	// demo1
	tel := voice.VoiceTel{Mobile: "1300000000", Nationcode: "86"}
	p := &voice.SendtVoiceParams{TplId: 404042, Params: []string{"1122", "3"}, Tel: tel}
	result, err := v.NewSendtVoice().Fetch(p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	// demo2
	result1, err1 := v.NewSendtVoice().Fetchs("13000000000", "86", 4356, []string{})
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(result1)
	}
}

// 语音验证码状态通知
func VoiceCodeNotify(w http.ResponseWriter, r *http.Request) {
	result, err := voice.VoiceCodeNotify(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 语音通知状态通知
func VoicePromptNotify(w http.ResponseWriter, r *http.Request) {
	result, err := voice.VoicePromptNotify(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 语音通知按键通知
func VoiceKeyNotify(w http.ResponseWriter, r *http.Request) {
	result, err := voice.VoiceKeyNotify(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 语音送达失败原因推送
func VoiceFailureNotify(w http.ResponseWriter, r *http.Request) {
	result, err := voice.VoiceFailureNotify(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 启动web服务器测试回调
func StartVoiceHttp() {
	http.HandleFunc("/voice_code", VoiceCodeNotify)
	http.HandleFunc("/voice_prompt", VoicePromptNotify)
	http.HandleFunc("/voice_key", VoiceKeyNotify)
	http.HandleFunc("/voice_failure", VoiceFailureNotify)
	// 如果在本地局域网则无法回调，可以放在服务器上执行回调或者使用 www.ngrok.cc 工具
	http.ListenAndServe("127.0.0.1:8888", nil)
}
