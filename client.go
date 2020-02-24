package baas_clientgosdk

const (
	loginPath  = "/v1/login"
	invokePath = "/v1/continvoke"
	queryPath  = "/v1/contquery"
)

type ClientInstance struct {
	Node   string
	Logged bool
	Token  string

	Acct   string
	PwdMd5 string
}

func (c *ClientInstance) Login(acct, pwd string) (ok bool, msg string, err error) {
	uri := c.Node + loginPath
	// pwd = md5Str(pwd)
	tk, msg, err := ApiLogin(uri, acct, pwd)
	if err != nil {
		return
	}
	if msg != "" {
		return
	}
	ok = true
	c.Logged = true
	c.Token = tk
	return
}

func (c *ClientInstance) Invoke(cont, method string, args map[string]string) (rst, msg string,
	err error) {
	uri := c.Node + invokePath
	return ApiInvoke(uri, c.Token, cont, method, args)
}

func (c *ClientInstance) Query(cont, method string, args map[string]string) (rst, msg string,
	err error) {
	uri := c.Node + queryPath
	return ApiQuery(uri, c.Token, cont, method, args)
}
