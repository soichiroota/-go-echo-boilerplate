package models

import (
	"sample/db"

	"context"
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:t"`

	ID        int64     `bun:"id,pk,autoincrement"`
	Email     string        `bun:"email,notnull"`
	Password  string        `bun:"password,notnull"`
	Token     string        `bun:"token`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	DeletedAt time.Time `bun:deleted_at",soft_delete,nullzero"`
}

func NewCreateUserTable() error {
	ctx := context.Background()
	_, err := db.DB.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
	return err
}

func GetAllUsers() ([]User, error) {
	var users []User
	ctx := context.Background()
	err := db.DB.NewSelect().Model(&users).Order("created_at").Scan(ctx)
	return users, err
}

func CreateUser(u *User) (*User, error) {
	ctx := context.Background()
	_, err := db.DB.NewInsert().Model(u).Exec(ctx)
	return u, err
}

func GetUserById(id int64) (*User, error) {
	var user User
	ctx := context.Background()
	err := db.DB.NewSelect().Model(&user).Where("id = ?", id).Scan(ctx)
	return &user, err
}

func getUserByEmailAndPassword(email string, password string) (*User, error) {
	var user User
	ctx := context.Background()
	err := db.DB.NewSelect().Model(&user).Where("email = ? and password = ?", email, password).Scan(ctx)
	return &user, err
}

func UpdateUser(u *User) (*User, error) {
	ctx := context.Background()
	var orig User
	orig.Email = u.Email
	orig.Password = u.Password
	orig.Token = u.Token
	_, err := db.DB.NewUpdate().Model(&orig).Where("id = ?", u.ID).Exec(ctx)
	return u, err
}