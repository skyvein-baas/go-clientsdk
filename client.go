package baas_clientgosdk

const (
	loginPath    = "/v1/login"
	getTokenPath = "/v1/gettoken"
	invokePath   = "/v1/continvoke"
	queryPath    = "/v1/contquery"
)

type ClientInstance struct {
	Node   string
	Logged bool
	Token  string

	Acc *Acct
}

func (c *ClientInstance) Login(acct, pwd string) (ok bool, accP *Acct, msg string, err error) {
	uri := c.Node + loginPath
	// pwd = md5Str(pwd)
	acc, msg, err := ApiLogin(uri, acct, pwd)
	if err != nil {
		return
	}
	if msg != "" {
		return
	}
	ok = true
	accP = &acc
	return
}

func (c *ClientInstance) GetToken(acc *Acct) (ok bool, msg string, err error) {
	if acc == nil {
		return
	}
	uri := c.Node + getTokenPath
	token, msg, err := ApiGettoken(uri, *acc)
	if err != nil {
		return
	}
	if msg != "" {
		return
	}
	c.Token = token
	c.Logged = true
	ok = true
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
