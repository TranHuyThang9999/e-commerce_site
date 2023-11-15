package model

// Role struct đại diện cho vai trò của người dùng
type Role struct {
	ID     int64 `json:"id"`
	Admin  int   `json:"admin"`
	Seller int   `json:"seller"`
	Buyer  int   `json:"buyer"`
}
