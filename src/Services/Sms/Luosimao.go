/**
 * Luosimao 短信发送封装。
 * @author fingerQin
 * @date 2018-12-05
 */

package Sms

import "log"

import "github.com/imroc/req"

type Luosimao struct {
}

func (s Luosimao) Send(mobile, content string) bool {
	return true
}

func (s Luosimao) request(mobile, content string) string {
	authHeader := req.Header{
		"Accept": "application/json",
		"user":   "api:key-a4957b61dce1f9329694030f3b0b25fe",
	}
	param := req.Param{
		"mobile":  mobile,
		"message": content,
	}
	apiUrl := "http://sms-api.luosimao.com/v1/send.json"
	req.Debug = true
	r, err := req.Post(apiUrl, authHeader, param)
	if err != nil {
		log.Fatal(err)
	}
	r.ToJSON(&foo)       // 响应体转成对象
	log.Printf("%+v", r) // 打印详细信息
	if r.error != 0 {
		return false
	} else {
		return true
	}
}
