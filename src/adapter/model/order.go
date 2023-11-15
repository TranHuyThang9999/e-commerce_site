package model

// Order struct đại diện cho thông tin đơn hàng
type Order struct {
	ID            int64 `json:"id"`
	IDProduct     int64 `json:"id_product"`
	IDBuyer       int64 `json:"id_buyer"`
	IDSeller      int64 `json:"id_seller"`
	OrderStatus   int   `json:"order_status"`
	PaymentStatus int   `json:"payment_status"`
	CreatedAt     int   `json:"created_at"`
	UpdatedAt     int   `json:"updated_at"`
}
