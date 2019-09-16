package template

import (
	"encoding/json"
	"fmt"
	"github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/context"
	"github.com/sunnyos/tencentSms/utils"
	"time"
)

type Template struct {
	Context *context.Context
}

const (
	ADD_TEMPLATE_URL = "https://yun.tim.qq.com/v5/tlssmssvr/add_template"
	MOD_TEMPLATE_URL = "https://yun.tim.qq.com/v5/tlssmssvr/mod_template"
	DEL_TEMPLATE_URL = "https://yun.tim.qq.com/v5/tlssmssvr/del_template"
	GET_TEMPLATE_URL = "https://yun.tim.qq.com/v5/tlssmssvr/get_template"
	ORDINARY_SMS     = 0
	MARKETING_SMS    = 1
)

type AddParams struct {
	Remark        string `json:"remark"`
	International int    `json:"international"`
	Sig           string `json:"sig"`
	Text          string `json:"text"`
	Time          int64  `json:"time"`
	Title         string `json:"title"`
	Type          int    `json:"type"`
}

type ModParams struct {
	Remark string `json:"remark"`
	Sig    string `json:"sig"`
	Text   string `json:"text"`
	Time   int64  `json:"time"`
	Title  string `json:"title"`
	TplId  int64  `json:"tpl_id"`
	Type   int    `json:"type"`
}

type DelParams struct {
	Sig   string `json:"sig"`
	Time  int64  `json:"time"`
	TplId []int  `json:"tpl_id"`
}

type GetParams struct {
	Sig   string `json:"sig"`
	Time  int64  `json:"time"`
	TplId []int  `json:"tpl_id"`
}

type GetPageParams struct {
	Sig     string  `json:"sig"`
	Time    int64   `json:"time"`
	TplPage TplPage `json:"tpl_page"`
}

type DelResult struct {
	Result int    `json:"result"`
	Errmsg string `json:"errmsg"`
}

type TplPage struct {
	Max    int `json:"max"`
	Offset int `json:"offset"`
}
type Result struct {
	Result int        `json:"result"`
	Errmsg string     `json:"errmsg"`
	Data   ResultData `json:"data"`
}

type ResultData struct {
	Id            int    `json:"id"`
	International int    `json:"international"`
	Status        int    `json:"status"`
	Text          string `json:"text"`
	Type          int    `json:"type"`
}

type GetResult struct {
	Result int             `json:"result"`
	Errmsg string          `json:"errmsg"`
	Total  int             `json:"total,omitempy"`
	Count  int             `json:"count,omitempy"`
	Data   []GetResultData `json:"data,omitempy"`
}

type GetResultData struct {
	Id            int    `json:"id"`
	International int    `json:"international"`
	Status        int    `json:"status"`
	Reply         string `json:"reply"`
	Text          string `json:"text"`
	Type          int    `json:"type"`
	Title         string `json:"title"`
	ApplyTime     string `json:"apply_time"`
	ReplyTime     string `json:"reply_time"`
}

func NewTemplate(cfg *config.Config) *Template {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Template{Context: context}
}

func copyConfigToContext(cfg *config.Config, context *context.Context) {
	context.AppId = cfg.AppId
	context.AppKey = cfg.AppKey
	context.Sign = cfg.Sign
}

func (t *Template) Add(p *AddParams) (*Result, error) {
	p.Time = time.Now().Unix()
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(t.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", ADD_TEMPLATE_URL, t.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &Result{}
	err = json.Unmarshal(res, result)
	return result, err
}

func (t *Template) Mod(p *ModParams) (*Result, error) {
	p.Time = time.Now().Unix()
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(t.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", MOD_TEMPLATE_URL, t.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &Result{}
	err = json.Unmarshal(res, result)
	return result, err
}

func (t *Template) Del(tplId []int) (*DelResult, error) {
	p := &DelParams{
		Time:  time.Now().Unix(),
		TplId: tplId,
	}
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(t.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", DEL_TEMPLATE_URL, t.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &DelResult{}
	err = json.Unmarshal(res, result)
	return result, err
}

func (t *Template) Get(tplId []int) (*GetResult, error) {
	p := &GetParams{
		Time:  time.Now().Unix(),
		TplId: tplId,
	}
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(t.Context.AppKey, p.Time, random)
	result, err := t.get(p, random)
	return result, err
}

func (t *Template) GetPage(max int, offset int) (*GetResult, error) {
	p := &GetPageParams{
		Time:    time.Now().Unix(),
		TplPage: TplPage{Max: max, Offset: offset},
	}
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(t.Context.AppKey, p.Time, random)
	result, err := t.get(p, random)
	return result, err
}

func (t *Template) get(p interface{}, random string) (*GetResult, error) {
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", GET_TEMPLATE_URL, t.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &GetResult{}
	err = json.Unmarshal(res, result)
	return result, err
}
