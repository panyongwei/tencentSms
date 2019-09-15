package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func GetSignatureWithMobile(appkey string, mobile string, time int64, random string) (string) {
	data := fmt.Sprintf("appkey=%s&random=%s&time=%d&mobile=%s", appkey, random, time, mobile)
	return getSignature(data)
}

func GetSignatureWithOutMobile(appkey string, time int64, random string) (string) {
	data := fmt.Sprintf("appkey=%s&random=%s&time=%d", appkey, random, time)
	return getSignature(data)
}

func getSignature(data string) (string) {
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}
