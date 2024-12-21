package model

type User struct {
	Name  string `json:"name" redis:"name"`
	Login string `json:"login" redis:"login"`
}
