package baas_clientgosdk

import (
	"fmt"
)

func NewInstance(node string) (cli ClientInstance) {
	cli = ClientInstance{
		Node: node,
	}
	return
}

func (c *ClientInstance) EnsureLogin(acct, pwd string) (ok bool, msg string, err error) {
	pwd = md5Str(pwd)
	ok, msg, err = c.Login(acct, pwd)
	if err != nil || msg != "" || !ok {
		return
	}
	c.Acct = acct
	c.PwdMd5 = pwd
	return
}

func (c *ClientInstance) EnsureInvoke(cont, method string, args map[string]string) (rst, msg string,
	err error) {
	if !c.Logged {
		err = fmt.Errorf("not logged")
		return
	}
	rst, msg, err = c.Invoke(cont, method, args)
	if err != nil || (msg != "" && msg != "token") {
		return
	}
	if msg == "token" {
		okl, msgl, errl := c.Login(c.Acct, c.PwdMd5)
		if !okl {
			c.Logged = false
			c.Token = ""
		}
		msg, err = msgl, errl
		return
	}
	return
}

func (c *ClientInstance) EnsureQuery(cont, method string, args map[string]string) (rst, msg string,
	err error) {
	if !c.Logged {
		err = fmt.Errorf("not logged")
		return
	}
	rst, msg, err = c.Query(cont, method, args)
	if err != nil || (msg != "" && msg != "token") {
		return
	}
	if msg == "token" {
		okl, msgl, errl := c.Login(c.Acct, c.PwdMd5)
		if !okl {
			c.Logged = false
			c.Token = ""
		}
		msg, err = msgl, errl
		return
	}
	return
}
