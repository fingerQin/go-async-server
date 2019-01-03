/**
 * Luosimao 短信发送封装。
 * @author fingerQin
 * @date 2018-12-05
 */

package Sms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	smsURL   = "http://sms-api.luosimao.com/v1/send.json"
	voiceURL = "http://voice-api.luosimao.com/v1/verify.json"
)

// Client : a SDK client for luosimao.com
type Luosimao struct {
	KeySMS   string
	KeyVoice string
}

// SendSMS : send a SMS message to the mobile number
func (c Luosimao) SendSMS(mobile interface{}, message string) error {
	params := url.Values{}
	params.Add("mobile", fmt.Sprint(mobile))
	params.Add("message", message)
	body, err := c.httpDo("POST", smsURL, c.KeySMS, params)
	if err != nil {
		return err
	}
	result := struct {
		Msg   string `json:"msg"`
		Error int    `json:"error"`
	}{"Unkonw Error", -1}
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.Error != 0 {
		err = fmt.Errorf(result.Msg)
	}
	return err
}

// SendVoice : send a voico message to the mobile number
func (c Luosimao) SendVoice(mobile, code interface{}) error {
	params := url.Values{}
	params.Add("mobile", fmt.Sprint(mobile))
	params.Add("code", fmt.Sprint(code))
	body, err := c.httpDo("POST", voiceURL, c.KeyVoice, params)
	if err != nil {
		return err
	}
	result := struct {
		Msg   string `json:"msg"`
		Error int    `json:"error"`
	}{"Unkown error.", -1}
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}
	if result.Error != 0 {
		err = fmt.Errorf(result.Msg)
	}
	return err
}

func (c Luosimao) httpDo(method, url, key string, params url.Values) (body []byte, err error) {
	httpClient := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest(method, url, strings.NewReader(params.Encode()))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("api", key)
	var resp *http.Response
	resp, err = httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
