package baas_clientgosdk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ApiLogin(uri string, acct, pwd string) (acc Acct, msg string, err error) {
	m := make(map[string]string)
	m["account"] = acct
	m["pwd"] = pwd
	d, _ := json.Marshal(m)
	r := bytes.NewReader(d)
	resp, err := http.Post(uri, "application/json;chartset=uft-8", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var rsp BaseRsp
	json.Unmarshal(body, &rsp)
	if rsp.Status != 1 {
		msg = rsp.Msg
		return
	}
	tkm, ok := rsp.Data[0].(map[string]interface{})
	if ok {
		d, _ := json.Marshal(tkm)
		json.Unmarshal(d, &acc)
	}
	return
}

type Acct struct {
	// 地址
	Address string `json:"address"`
	// 私钥
	Prikey string `json:"prikey"`
	// 共钥
	Pubkey string `json:"pubkey"`
	// 助记词
	Mnemonic string `json:"mnemonic"`
}

func ApiGettoken(uri string, acc Acct) (token, msg string, err error) {
	d, _ := json.Marshal(acc)
	r := bytes.NewReader(d)
	resp, err := http.Post(uri, "application/json;chartset=uft-8", r)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var rsp BaseRsp
	json.Unmarshal(body, &rsp)
	if rsp.Status != 1 {
		msg = rsp.Msg
		return
	}
	tkm, ok := rsp.Data[0].(map[string]interface{})
	if ok {
		token = tkm["token"].(string)
	}
	return
}

func ApiInvoke(uri, token string, cont, method string, args map[string]string) (
	rst, msg string, err error) {
	m := make(map[string]interface{})
	m["contract"] = cont
	m["method"] = method
	m["args"] = args
	d, _ := json.Marshal(m)
	r := bytes.NewReader(d)
	cli := http.Client{}
	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;chartset=uft-8")
	req.Header.Set("Token", token)
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var rsp BaseRsp
	json.Unmarshal(body, &rsp)
	if rsp.Status != 1 {
		msg = rsp.Msg
		return
	}
	rst, _ = rsp.Data[0].(string)
	return
}

func ApiQuery(uri, token string, cont, method string, args map[string]string) (
	rst, msg string, err error) {
	m := make(map[string]interface{})
	m["contract"] = cont
	m["method"] = method
	m["args"] = args
	d, _ := json.Marshal(m)
	r := bytes.NewReader(d)
	cli := http.Client{}
	req, err := http.NewRequest("POST", uri, r)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json;chartset=uft-8")
	req.Header.Set("Token", token)
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var rsp BaseRsp
	json.Unmarshal(body, &rsp)
	if rsp.Status != 1 {
		msg = rsp.Msg
		return
	}
	rst, _ = rsp.Data[0].(string)
	return
}
