package examples

import (
	"fmt"
	"github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/sms"
	"net/http"
	"time"
)

// 指定模板单发短信
func TestSingleSendSms() {
	c := &config.Config{AppId: "", AppKey: "", Sign: ""}
	s := sms.NewSms(c)

	// demo1
	tel := sms.SmsTel{Mobile: "1300000000", Nationcode: "86"}
	p := &sms.Params{Params: []string{"1234", "5"}, Tel: tel, TplId: 420450}
	res, err := s.GetSmsSender().Fetch(p)
	if err != nil {

	} else {
		fmt.Println(res)
	}

	// demo2
	res1, err1 := s.GetSmsSender().Fetchs("1300000000", "86", []string{"33221", "5"}, 420450)
	if err1 != nil {

	} else {
		fmt.Println(res1)
	}
}

// 指定模板群发短信
func TestMultiSms() {
	c := &config.Config{AppId: "", AppKey: "", Sign: ""}
	s := sms.NewSms(c)

	// demo1
	tel := []sms.SmsTel{sms.SmsTel{Mobile: "130000000", Nationcode: "86"}, sms.SmsTel{Mobile: "1300000001", Nationcode: "86"}}
	p := &sms.MultiParams{Params: []string{"1234", "5"}, Tel: tel, TplId: 420450}
	res, err := s.GetSendMultiSms().Fetch(p)
	if err != nil {

	} else {
		fmt.Println(res)
	}

	// demo2
	res1, err1 := s.GetSendMultiSms().Fetchs([]string{"1300000000", "1300000001"}, []string{"86", "86"}, []string{"33221", "5"}, 420450)
	if err1 != nil {

	} else {
		fmt.Println(res1)
	}
}

// 拉取单个手机短信状态
func PullStatus() {
	c := &config.Config{AppId: "", AppKey: ""}
	s := sms.NewSms(c)
	p := &sms.PullStatusParams{
		BeginTime:  time.Now().Unix() - (3600 * 24),
		EndTime:    time.Now().Unix(),
		Max:        100,
		Mobile:     "13000000000",
		Nationcode: "86",
		Type:       0,
	}
	res, err := s.GetPullStatus().Fetch(p)
	if err != nil {

	} else {
		fmt.Println(res)
	}
}

// 拉取短信状态
func PullMultiStatus() {
	c := &config.Config{AppId: "", AppKey: ""}
	s := sms.NewSms(c)
	p := &sms.PullMultiStatusParams{
		Max:  100,
		Type: 0,
	}
	res, err := s.GetPullMultiStatus().Fetch(p)
	if err != nil {

	} else {
		fmt.Println(res)
	}
}

// 短信状态回调
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	nofity, err := sms.Notify(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(nofity)
	}
}

// 短信回复回调
func Reply(w http.ResponseWriter, r *http.Request) {
	reply, err := sms.Reply(r)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reply)
		fmt.Println(reply.Text)
	}

}

// 启动web服务器测试回调
func StartHttp() {
	http.HandleFunc("/call", Reply)
	http.HandleFunc("/status", StatusHandler)
	// 如果在本地局域网则无法回调，可以放在服务器上执行回调或者使用 www.ngrok.cc 工具
	http.ListenAndServe("127.0.0.1:8888", nil)
}
