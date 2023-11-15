package entities

type UsersReq struct {
	Name    string `form:"name"`
	Age     int    `form:"age"`
	Address string `form:"address"`
}
type UsersResp struct {
	Id     int64  `json:"id"`
	Result Result `json:"result"`
}
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
