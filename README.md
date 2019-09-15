# 腾讯短信 GO SDK

官网没有 go 版本的 sdk 自己着手写了一份简单的 sdk，大家可以直接安装了使用。代码都是开源的，可以直接开箱使用。

### 获取sdk

```shell
go get github.com/sunnyos/tencentSms
```

### 基本配置

```go
config := &sms.Config{
	AppId: "SDK AppID",
	AppKey: "App Key",
	Sign: "短信签名"
}
```

参数  | 说明
------------- | -------------
AppId|应用信息里的SDK AppID
AppKey|应用信息里的App Key
Sign|短信签名

## 短信API

### 指定模板单发短信

```go
s := sms.NewSms(config)

// demo1
tel := sms.SmsTel{Mobile: "1300000000", Nationcode: "86"}
p := &sms.Params{Params: []string{"1234", "5"}, Tel: tel, TplId: 420450}
res, err := s.GetSmsSender().Fetch(p)

// demo2
res1, err1 := s.GetSmsSender().Fetchs("1300000000", "86", []string{"33221", "5"}, 420450)
```

### 指定模板群发短信

```go
s := sms.NewSms(config)

// demo1
tel := []sms.SmsTel{sms.SmsTel{Mobile: "130000000", Nationcode: "86"}, sms.SmsTel{Mobile: "1300000001", Nationcode: "86"}}
p := &sms.MultiParams{Params: []string{"1234", "5"}, Tel: tel, TplId: 420450}
res, err := s.GetSendMultiSms().Fetch(p)

// demo2
res1, err1 := s.GetSendMultiSms().Fetchs([]string{"1300000000", "1300000001"}, []string{"86", "86"}, []string{"33221", "5"}, 420450)
```

### 拉取单个手机短信状态

```go
s := sms.NewSms(config)
p := &sms.PullStatusParams{
	BeginTime:  time.Now().Unix() - (3600 * 24),
	EndTime:    time.Now().Unix(),
	Max:        100,
	Mobile:     "13000000000",
	Nationcode: "86",
	Type:       0,
}
res, err := s.GetPullStatus().Fetch(p)
```

### 拉取短信状态

```go
s := sms.NewSms(config)
p := &sms.PullMultiStatusParams{
	Max:  100,
	Type: 0,
}
res, err := s.GetPullMultiStatus().Fetch(p)
```

### 短信状态回调

```go
nofity, err := sms.Notify(r)
```

r 是 *http.Request

### 短信回复回调

```go
reply, err := sms.Reply(r)
```

r 是 *http.Request

## 语音API

### 发送语音验证码

```go
v := voice.NewVoice(config)

// demo1
tel := voice.VoiceTel{Mobile: "1300000000", Nationcode: "86"}
p := &voice.SendcVoiceParams{Msg: "1122", Tel: tel}
result, err := v.NewSendcVoice().Fetch(p)

// demo2
result1, err1 := v.NewSendcVoice().Fetchs("1233", "13000000000", "86")
```

### 指定模板发送语音通知

```go
v := voice.NewVoice(config)

// demo1
tel := voice.VoiceTel{Mobile: "1300000000", Nationcode: "86"}
p := &voice.SendtVoiceParams{TplId: 404042, Params: []string{"1122", "3"}, Tel: tel}
result, err := v.NewSendtVoice().Fetch(p)

// demo2
result1, err1 := v.NewSendtVoice().Fetchs("13000000000", "86", 4356, []string{})
```

### 语音验证码状态通知

```go
result, err := voice.VoiceCodeNotify(r)
```

r 是 *http.Request

### 语音通知状态通知

```go
result, err := voice.VoicePromptNotify(r)
```

r 是 *http.Request

### 语音通知按键通知

```go
result, err := voice.VoiceKeyNotify(r)
```

r 是 *http.Request

### 语音送达失败原因推送

```go
result, err := voice.VoiceFailureNotify(r)
```

r 是 *http.Request




