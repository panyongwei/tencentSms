package sign

import (
	"encoding/json"
	"fmt"
	"github.com/sunnyos/tencentSms/config"
	"github.com/sunnyos/tencentSms/context"
	"github.com/sunnyos/tencentSms/utils"
	"time"
)

const (
	ADD_SIGN_URL = "https://yun.tim.qq.com/v5/tlssmssvr/add_sign"
	MOD_SIGN_URL = "https://yun.tim.qq.com/v5/tlssmssvr/mod_sign"
	DEL_SIGN_URL = "https://yun.tim.qq.com/v5/tlssmssvr/del_sign"
	GET_SIGN_URL = "https://yun.tim.qq.com/v5/tlssmssvr/get_sign"
)

type Sign struct {
	Context *context.Context
}

type SignParams struct {
	Pic           string `json:"pic"`
	International int    `json:"international,omitempy"`
	Remark        string `json:"remark"`
	Sig           string `json:"sig"`
	Text          string `json:"text"`
	Time          int64  `json:"time"`
	SignId        int64  `json:"sign_id,omitempy"`
}

type DelGetSignParams struct {
	Sig      string        `json:"sig"`
	SignId   []int64       `json:"sign_id,omitempy"`
	Time     int64         `json:"time"`
	SignPage GetPageResult `json:"sign_page,omitempy"`
}

type SignResult struct {
	Result int            `json:"result"`
	Errmsg string         `json:"errmsg"`
	Msg    string         `json:"msg,omitempy"`
	Data   SignResultData `json:"data,omitempy"`
}

type SignResultData struct {
	Id            int    `json:"id"`
	International int    `json:"international"`
	Status        int    `json:"status"`
	Text          string `json:"text"`
}

type GetPageResult struct {
	Max    int `json:"max"`
	Offset int `json:"offset"`
}

type GetResult struct {
	Result int             `json:"result"`
	Errmsg string          `json:"errmsg"`
	Msg    string          `json:"msg"`
	Total  int             `json:"total"`
	Count  int             `json:"count,omitempy"`
	Data   []GetResultData `json:"data,omitempy"`
}

type GetResultData struct {
	Id            int    `json:"id"`
	International int    `json:"international"`
	Reply         string `json:"reply"`
	Status        int    `json:"status"`
	Text          string `json:"text"`
	ApplyTime     string `json:"apply_time"`
	ReplyTime     string `json:"reply_time"`
}

func NewSign(cfg *config.Config) *Sign {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Sign{context}
}

func copyConfigToContext(cfg *config.Config, context *context.Context) {
	context.AppId = cfg.AppId
	context.AppKey = cfg.AppKey
	context.Sign = cfg.Sign
}

func (s *Sign) Add(p *SignParams) (*SignResult, error) {
	p.Time = time.Now().Unix()
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(s.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", ADD_SIGN_URL, s.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &SignResult{}
	err = json.Unmarshal(res, result)
	return result, err
}

func (s *Sign) Mod(p *SignParams) (*SignResult, error) {
	p.Time = time.Now().Unix()
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(s.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", MOD_SIGN_URL, s.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &SignResult{}
	err = json.Unmarshal(res, result)
	return result, err
}

func (s *Sign) Del(signId []int64) (*SignResult, error) {
	p := &DelGetSignParams{
		SignId: signId,
		Time:   time.Now().Unix(),
	}
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(s.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", DEL_SIGN_URL, s.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &SignResult{}
	err = json.Unmarshal(res, result)
	return result, err
}

func (s *Sign) Get(signId []int64) (*GetResult, error) {
	p := &DelGetSignParams{
		SignId: signId,
		Time:   time.Now().Unix(),
	}
	result, err := s.get(p)
	return result, err
}

func (s *Sign) GetPage(max int, offset int) (*GetResult, error) {
	signId := []int64{}
	p := &DelGetSignParams{
		SignId:   signId,
		SignPage: GetPageResult{Max: max, Offset: offset},
		Time:     time.Now().Unix(),
	}

	result, err := s.get(p)
	return result, err
}

func (s *Sign) get(p *DelGetSignParams) (*GetResult, error) {
	random := utils.CreateRandom()
	p.Sig = utils.GetSignatureWithOutMobile(s.Context.AppKey, p.Time, random)
	uri := fmt.Sprintf("%s?sdkappid=%s&random=%s", GET_SIGN_URL, s.Context.AppId, random)
	res, err := utils.PostJSON(uri, p)
	if err != nil {
		return nil, err
	}
	result := &GetResult{}
	err = json.Unmarshal(res, result)
	return result, err
}
