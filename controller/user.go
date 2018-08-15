package controller

import "net/http"

type User struct{}

type UserModel struct {
	Name string
}

func (u *User) Get(w http.ResponseWriter, r *http.Request) error {
	users := []*UserModel{
		&UserModel{Name: "hoge"},
	}

	return JSON(w, 200, users)
}
