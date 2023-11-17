package enums

const (
	SUCCESS_CODE = 0
	SUCCESS_MESS = "Success"

	LOGIN_ERR_CODE = 2
	LOGIN_ERR_MESS = "Login failed"

	DB_ERR_CODE = 4
	DB_ERR_MESS = "Database error"

	TRANSACTION_INVALID_CODE = 6
	TRANSACTION_INVALID_MESS = "invalid transaction"

	ACCOUNT_NOT_EXIST_CODE = 8
	ACCOUNT_NOT_EXIST_MESS = "ACCOUNT  not exist"

	ACCOUNT_EXIST_CODE = 10
	ACCOUNT_EXIST_MESS = "ACCOUNT  exist"

	HASH_PASSWORD_ERR      = 12
	HASH_PASSWORD_ERR_MESS = "error hash password"

	SEND_EMAIL_CODE_ERROR = 14
	SEND_EMAIL_MESS_ERROR = "error send email author 3"

	EMAIL_ACCOUNT_EXITS_CODE = 16
	EMAIL_ACCOUNT_EXITS_MESS = "email exits"

	PHONE_NUMBER_EXITS_CODE = 18
	PHONE_NUMBER_EXITS_MESS = "phone number exits"

	ERROR_SAVE_IMAGE_CODE  = 20
	ERRORL_SAVE_IMAGE_MESS = "error save image"

	STORE_NAME_EXITS_CODE = 22
	STORE_NAME_EXITS_MESS = "StoreName exits"

	CREATE_TOKEN      = 22
	CREATE_TOKEN_MESS = "create token error"

	VERIFIED_ACCOUNT_ERROR_CODE = 24
	VERIFIED_ACCOUNT_ERROR_MESS = "verified code error"

	CONVERT_TO_NUMBER_CODE = 26
	CONVERT_TO_NUMBER_MESS = "error convert string to number"

	ACCOUNT_NOT_VERIFIED_CODE = 28
	ACCOUNT_NOT_VERIFIED_MESS = "account not verified"

	ACCOUNT_OR_PASSWORD_WRONG_CODE = 30
	ACCOUNT_OR_PASSWORD_WRONG_MESS = "Wrong account or password"

	PRODUCT_EMPTY_CODE = 32
	PRODUCT_EMPTY_MESS = "product empty"
)

const (
	ROLE_ADMIN     = 3
	ROLE_NOT_ADMIN = 5

	ROLE_USER_BUYER_ACTIVE = 7
	ROLE_USER_BUYER_LOCK   = 9

	ROLE_USER_SELLER_ACTIVE = 11
	ROLE_USER_SELLER_LOCK   = 13

	IS_VERIFIED  = 15
	NOT_VERIFIED = 17
)
