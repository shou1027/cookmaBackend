package model

import (
	"errors"
	"net/mail"
	"unicode/utf8"
)

const (
	nameLengthMax = 20
	nameLengthMin = 1
)

type User struct {
	id       int64
	name     string
	email    string
	password string
}

func Reconstruct(
	id int64,
	name string,
	email string,
	password string,
) (*User, error) {
	return newUser(
		id,
		name,
		email,
		password,
	)
}

func NewUser(
	name string,
	email string,
	password string,
) (*User, error) {
	return newUser(
		0,
		name,
		email,
		password,
	)
}

func newUser(
	id int64,
	name string,
	email string,
	password string,
) (*User, error) {
	// 名前バリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errors.New("名前の文字数が不正です。")
	}

	// メールアドレスバリデーション
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errors.New("メールアドレスが不正です。")
	}

	//パスワード
	//TODO バリデーションが必要か検討

	return &User{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}, nil
}

// ユーザーIDを返却
func (u User) GetId() int64 {
	return u.id
}

// ユーザー名を返却
func (u User) GetName() string {
	return u.name
}

// メールアドレスを返却
func (u User) GetEmail() string {
	return u.email
}

// パスワードを返却
func (u User) GetPassword() string {
	return u.password
}
