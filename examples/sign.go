package examples

import (
	"fmt"
	"github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/sign"
)

// 添加签名
func TestAddSign() {
	c := &config.Config{AppId: "", AppKey: ""}
	s := sign.NewSign(c)
	p := &sign.SignParams{Text: "测试"}
	result, err := s.Add(p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 修改签名
func TestModSign() {
	c := &config.Config{AppId: "", AppKey: ""}
	s := sign.NewSign(c)
	p := &sign.SignParams{Text: "测试", SignId: 1234}
	result, err := s.Mod(p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 删除签名
func TestDelSign() {
	c := &config.Config{AppId: "", AppKey: ""}
	s := sign.NewSign(c)
	signId := []int64{235741,111}
	result, err := s.Del(signId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 指定签名ID拉取签名状态
func TestGetSign() {
	c := &config.Config{AppId: "", AppKey: ""}
	s := sign.NewSign(c)
	signId := []int64{235741,111}
	result, err := s.Get(signId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 分页全量拉取签名状态
func TestGetPageSign() {
	c := &config.Config{AppId: "", AppKey: ""}
	s := sign.NewSign(c)
	result, err := s.GetPage(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
