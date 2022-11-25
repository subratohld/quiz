package models

type AuthData struct{}

type User struct {
	tableName    struct{} `pg:"tbl_user,discard_unknown_columns"`
	ID           int64    `pg:",pk"`
	UserID       string
	EmailID      string
	MobileNumber string
	FirstName    string
	LastName     string
	Address      string
}
