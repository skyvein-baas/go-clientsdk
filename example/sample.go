package main

import (
	"log"

	baassdk "git.skyvein.net/service/baas_clientgosdk"
)

const (
	node = "http://122.224.183.34:31353"
	acct = "13255554444"
	pwd  = "123456"
)

func main() {
	baasCli := baassdk.NewInstance(node)

	ok, msg, err := baasCli.EnsureLogin(acct, pwd)
	if !ok {
		log.Println(msg, err)
		return
	}
	// 先确认使用的账号有gas
	log.Println(baasCli.EnsureInvoke("mycounter", "increase", map[string]string{"key": "t1"}))
	log.Println(baasCli.EnsureQuery("mycounter", "get", map[string]string{"key": "t1"}))
}
