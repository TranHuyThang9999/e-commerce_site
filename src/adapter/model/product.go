package model

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

// ProductType struct đại diện cho loại sản phẩm
type ProductType struct {
	ID int64 `json:"id"`
}
