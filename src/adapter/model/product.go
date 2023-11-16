package model

import "mime/multipart"

// Product struct đại diện cho thông tin sản phẩm
type Product struct {
	ID            int64   `json:"id"`
	IDUser        int64   `json:"id_user"`
	NameProduct   string  `json:"name_product"`
	Quantity      int     `json:"quantity"`
	SellStatus    int     `json:"sell_status"`
	Price         float64 `json:"price"`
	Discount      float64 `json:"discount"`
	Manufacturer  string  `json:"manufacturer"`
	CreatedAt     int     `json:"created_at"`
	UpdatedAt     int     `json:"updated_at"`
	Describe      string  `json:"describe"`
	IDTypeProduct int64   `json:"id_type_product"`
}

type ProductReqCreate struct {
	UserName      string                  `form:"user_name"` //unique
	NameProduct   string                  `form:"name_product"`
	Quantity      int                     `form:"quantity"`
	SellStatus    int                     `form:"sell_status"`
	Price         float64                 `form:"price"`
	Discount      float64                 `form:"discount"`
	Manufacturer  string                  `form:"manufacturer"`
	Describe      string                  `form:"describe"`
	IDTypeProduct int64                   `form:"id_type_product"`
	Files         []*multipart.FileHeader `form:"files"`
}

type ProductRespCreate struct {
	Id     int64  `json:"id"`
	Result Result `json:"result"`
}

type ProductReqFindByForm struct {
	ID            int64   `form:"id"`
	IDUser        int64   `json:"id_user"`
	NameProduct   string  `form:"name_product"`
	Quantity      int     `form:"quantity"`
	SellStatus    int     `form:"sell_status"`
	Price         float64 `form:"price"`
	Discount      float64 `form:"discount"`
	Manufacturer  string  `form:"manufacturer"` // Nhà sản xuất
	CreatedAt     int     `form:"created_at"`
	UpdatedAt     int     `form:"updated_at"`
	Describe      string  `form:"describe"`
	IDTypeProduct int64   `form:"id_type_product"`
}

type ProductRespFindByForm struct {
	ID            int64   `json:"id"`
	NameProduct   string  `json:"name_product"`
	Quantity      int     `json:"quantity"`
	SellStatus    int     `json:"sell_status"`
	Price         float64 `json:"price"`
	Discount      float64 `json:"discount"`
	Manufacturer  string  `json:"manufacturer"`
	CreatedAt     int     `json:"created_at"`
	UpdatedAt     int     `json:"updated_at"`
	Describe      string  `json:"describe"`
	IDTypeProduct int64   `json:"id_type_product"`
}

type ProductsRespFindByForm struct {
	Result   Result                   `json:"result"`
	Total    int                      `json:"total"`
	Products []*ProductRespFindByForm `json:"products"`
}

// ProductType struct đại diện cho loại sản phẩm
type ProductType struct {
	ID int64 `json:"id"`
}
