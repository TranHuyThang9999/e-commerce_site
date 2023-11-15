package model

import "mime/multipart"

// Account struct đại diện cho thông tin của người dùng
type Account struct {
	ID          int64  `json:"id"`
	IDRole      int64  `json:"id_role"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Age         int    `json:"age"`
	Address     string `json:"address"`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`        //unique
	PhoneNumber string `json:"phone_number"` // unique
	UserName    string `json:"user_name"`    //unique
	Password    string `json:"password"`
	OTPCode     int64  `json:"otp_code"`
	OTPExpiry   int    `json:"otp_expiry"`
	IsVerified  int    `json:"is_verified"`
	StoreName   string `json:"store_name"`
	Notes       string `json:"notes"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
	Avatar      string `json:"avatar"`
}

type AccountReqCreate struct {
	FirstName   string                `form:"first_name" validate:"required"`
	LastName    string                `form:"last_name" validate:"required"`
	Age         int                   `form:"age" validate:"required"`
	Address     string                `form:"address" validate:"required"`
	Gender      int                   `form:"gender" validate:"required,oneof= 39 41 43"`
	Email       string                `form:"email" validate:"required"`
	PhoneNumber string                `form:"phone_number" validate:"required"`
	UserName    string                `form:"user_name" validate:"required"`
	Password    string                `form:"password" validate:"required"`
	StoreName   string                `form:"store_name" validate:"required"`
	Notes       string                `form:"notes" validate:"required"`
	File        *multipart.FileHeader `form:"file"`
}

type AccountRespCreate struct {
	Id     int64  `json:"id"`
	Result Result `json:"result"`
}

// ShippingAddress struct đại diện cho địa chỉ vận chuyển của người dùng
type ShippingAddress struct {
	ID          int64  `json:"id"`
	IDUser      int64  `json:"id_user"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Province    string `json:"province"`
	District    string `json:"district"`
	Commune     string `json:"commune"`
	Village     string `json:"village"`
	StreetName  string `json:"street_name"`
	Notes       string `json:"notes"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}
type AccountReqFindByForm struct {
	ID          int64  `form:"id"`
	FirstName   string `form:"first_name"`
	LastName    string `form:"last_name"`
	Age         int    `form:"age"`
	Address     string `form:"address"`
	Gender      int    `form:"gender"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phone_number"`
	UserName    string `form:"user_name"`
	StoreName   string `form:"store_name"`
	CreatedAt   int    `form:"created_at"`
	UpdatedAt   int    `form:"updated_at"`
}
type AccountRespFindByForm struct {
	Result   Result     `json:"result"`
	Total    int        `json:"total"`
	Accounts []*Account `json:"accounts"`
}

type AccountRespUpdateById struct {
	Id     int64  `json:"id"`
	Result Result `json:"result"`
}

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
