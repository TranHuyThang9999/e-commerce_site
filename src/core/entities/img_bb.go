package entities

import "mime/multipart"

type Data struct {
	URL string `json:"url"`
}

type ImagesReq struct {
	Name    string                  `form:"name"`
	Age     int                     `form:"age"`
	Address string                  `form:"address"`
	Files   []*multipart.FileHeader `form:"files"`
}
type InfomationImages struct {
	Id     int64  `json:"id"`
	Url    string `json:"url"`
	IdUser int64  `json:"id_user"`
}
