package users

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	GetUsers(user DTJson) ([]User, error)
	GetTotalUser(user DTJson) (int, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetUsers(user DTJson) ([]User, error) {
	var users []User
	sql := "Select * from users where 1=1 "

	if s := user.Search.Value; s != "" {
		sql = fmt.Sprintf("%s AND name like '%%%s%%' ", sql, s)
	}

	length := user.Length
	start := user.Start

	sql = fmt.Sprintf("%s LIMIT %d , %d", sql, start, length)

	fmt.Println(sql)

	err := r.db.Raw(sql).Scan(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) GetTotalUser(user DTJson) (int, error) {
	var users []User
	sql := "Select * from users where 1=1 "

	if s := user.Search.Value; s != "" {
		sql = fmt.Sprintf("%s AND (name like '%%%s%%'  OR age like '%%%s%%')", sql, s, s)
	}
	log.Println("Debug Query: " + sql)

	err := r.db.Raw(sql).Scan(&users).Error
	if err != nil {
		return len(users), err
	}

	return len(users), nil
}
