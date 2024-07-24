package model

import (
	"errors"
	"net/mail"
	"unicode/utf8"
)

const (
	nameLengthMax = 255
	nameLengthMin = 1
)

type User struct {
	id       int64
	name     string
	email    string
	password string
}

func NewUser(
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
		id:    id,
		name:  name,
		email: email,
	}, nil
}
