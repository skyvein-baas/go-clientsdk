package main

import (
	"log"

	baassdk "github.com/skyvein-baas/go-clientsdk"
)

const (
	node = "http://[$ip:$port]"
	acct = "xxxx"
	pwd  = "xxxx"
)

func main() {
	baasCli := baassdk.NewInstance(node)

	ok, msg, err := baasCli.EnsureLogin(acct, pwd)
	if !ok {
		log.Println(msg, err)
		return
	}
	// 也可以通过下载到的密钥登录获取cli
	// acc := &baassdk.Acct{
	// 	Address:  "",
	// 	Prikey:   "",
	// 	Pubkey:   "",
	// 	Mnemonic: "",
	// }
	// ok, msg, err = baasCli.GetToken(acc)
	// if err != nil || msg != "" || !ok {
	// log.Println(msg, err)
	// 	return
	// }
	// baasCli.Acc = acc

	// 先确认使用的账号有gas
	log.Println(baasCli.EnsureInvoke("counter1", "increase", map[string]string{"key": "t1"}))
	log.Println(baasCli.EnsureQuery("counter1", "get", map[string]string{"key": "t1"}))
}
