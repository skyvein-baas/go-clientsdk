#### baas项目的node节点调用sdk - golang

---

这是一个方便用户调用天脉baas节点的client项目，提供了节点的几个基础的方法调用。接口当前基于http请求，用户也可以自行封装。



### 快速开始

---

#### 环境配置

	- 版本控制工具：Git
	- 开发语言：Go 1.12.x及以上

#### 构建和运行例子

​	```git clone https://github.com/skyvein-baas/go-clientsdk.git```

​	```go run example/sample.go```



#### 具体方法

```js
	import (
		baassdk "github.com/skyvein-baas/go-clientsdk"
	)

	...

	// 构建client实例
	node := "http://[ip:port]"
	baasCli := baassdk.NewInstance(node)
	// 登录
	ok, msg, err := baasCli.EnsureLogin(acct, pwd)
	if !ok {
		log.Println(msg, err)
		return
	}

	// 先确认使用的账号有gas
	// 合约调用
	rst, msg, err := baasCli.EnsureInvoke("mycounter", "increase", map[string]string{"key": "t1"})
	if rst != "" {
		log.Println(msg, err)
		return
	}
	rst, msg, err = baasCli.EnsureQuery("mycounter", "get", map[string]string{"key": "t1"})
	if rst != "" {
		log.Println(msg, err)
		return
	}
```
