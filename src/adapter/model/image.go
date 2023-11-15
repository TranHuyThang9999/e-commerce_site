package model

// ImageStorage struct đại diện cho lưu trữ hình ảnh của sản phẩm
type ImageStorage struct {
	ID        int64  `json:"id"`
	Url       string `json:"url"`
	IDUser    int64  `json:"id_user"`
	IDProduct int64  `json:"id_product"`
}
type Data struct {
	URL string `json:"url"`
}
