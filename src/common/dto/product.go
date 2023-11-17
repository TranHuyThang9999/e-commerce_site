package dto

type ProductReqFindByForm struct {
	ID            int64   `form:"id"`
	UserName      string  `form:"user_name" validate:"required"`
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
	Limit         int     `form:"limit"`
	Offset        int     `form:"offset"`
}
