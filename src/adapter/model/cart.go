package model

// Cart struct đại diện cho giỏ hàng của người dùng
type Cart struct {
	ID        int64 `json:"id"`
	IDProduct int64 `json:"id_product"`
	IDSeller  int64 `json:"id_seller"`
	CreatedAt int   `json:"created_at"`
	UpdatedAt int   `json:"updated_at"`
}
