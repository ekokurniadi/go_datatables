package users

import (
	"fmt"
	"net/url"
)

type UserFormatter struct {
	Number int    `json:"no"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Link   string `json:"link"`
}

func FormatUser(user User) UserFormatter {
	u, _ := url.Parse("http://localhost:8080")
	userFormatter := UserFormatter{}
	userFormatter.Number = user.Number + 1
	userFormatter.Name = user.Name
	userFormatter.Age = user.Age
	userFormatter.Link = fmt.Sprintf("<a href='http://"+u.Host+"/api/v1/view/%d' data-color='#265ed7' style='color: rgb(38, 94, 215);'><i class='icon-copy dw dw-eye'></i></a> <a href='http://"+u.Host+"/api/v1/users/edit/%d' orange'><i class='icon-copy dw dw-edit1'></i></a> <a href='http://"+u.Host+"/api/v1/users/%d' data-color='#e95959' style='color: rgb(233, 89, 89);' onclick='javasciprt: return confirm(\"Are You Sure ?\")''><i class='icon-copy dw dw-delete-3'></i></a>", user.ID, user.ID, user.ID)

	return userFormatter
}

func FormatUsers(users []User) []UserFormatter {

	userFormatters := []UserFormatter{}

	for _, user := range users {
		userFormatter := FormatUser(user)
		userFormatters = append(userFormatters, userFormatter)
	}

	return userFormatters
}
