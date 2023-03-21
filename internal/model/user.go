package model

type User struct {
	Id           int
	Username     string
	PasswordHash string `db:"password_hash"`
}
