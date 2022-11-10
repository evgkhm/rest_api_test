package rest_api_test

type User struct {
	Id      int     `json:"id" db:"id"`
	Balance float64 `json:"balance" binding:"required"`
}
