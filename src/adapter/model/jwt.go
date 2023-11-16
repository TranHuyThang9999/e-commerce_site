package model

import "github.com/golang-jwt/jwt/v4"

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	AtExpires    int64  `json:"at_expires"`
	RefreshToken string `json:"refresh_token"`
	RtExpires    int64  `json:"rt_expires"`
}

type UserJwtClaim struct {
	*jwt.StandardClaims
	Id       int64  `json:"id"`
	UserName string `json:"user_name"`
	Role     int64  `json:"role"`
}
type ResponseLogin struct {
	Result   Result    `json:"result"`
	JwtToken *JwtToken `json:"jwt_token"`
}
type LoginReq struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
}
type VerifiedAccountResp struct {
	Result Result `json:"result"`
}
type ResendOtpResp struct {
	Result Result `json:"result"`
}
