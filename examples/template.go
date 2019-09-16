package examples

import (
	"fmt"
	"github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/template"
)

// 添加模版
func TestAddTemplate() {
	c := &config.Config{AppId: "", AppKey: ""}
	t := template.NewTemplate(c)

	params := &template.AddParams{Text: "测试添加模版{1}，api测试", Type: template.ORDINARY_SMS}
	result, err := t.Add(params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.Errmsg)
	}
}

// 修改模版
func TestModTemplate(){
	c := &config.Config{AppId: "", AppKey: ""}
	t := template.NewTemplate(c)

	params := &template.ModParams{Text: "测试添加模版{1}，api测试", Type: template.ORDINARY_SMS,TplId:123}
	result, err := t.Mod(params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 删除模版
func TestDelTemplate(){
	c := &config.Config{AppId: "", AppKey: ""}
	t := template.NewTemplate(c)

	result, err := t.Del([]int{1,2})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 指定模板 ID 拉取模版状态
func TestGetIdTemplateStatus(){
	c := &config.Config{AppId: "", AppKey: ""}
	t := template.NewTemplate(c)

	result, err := t.Get([]int{1677})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// 分页全量拉取
func TestGetPageTemplateStatus()  {
	c := &config.Config{AppId: "", AppKey: ""}
	t := template.NewTemplate(c)

	result, err := t.GetPage(1,0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}