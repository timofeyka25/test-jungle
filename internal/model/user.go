package model

type User struct {
	Id           int64
	Username     string
	PasswordHash string `db:"password_hash"`
}
