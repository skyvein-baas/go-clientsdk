package baas_clientgosdk

type BaseRsp struct {
	Status int           `json:"status"` // 1此次请求成功 0此次请求失败
	Msg    string        `json:"msg"`    // 失败原因
	Data   []interface{} `json:"data"`   // 此次请求成功下的接口响应
}
