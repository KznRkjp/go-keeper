package models

import "time"

//  Пользователь
type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// Хранимые данные пользователя - логин и пароль
type LoginPassword struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// Хранимые данные пользователя - банковская карта
type BankCard struct {
	ID             int64     `json:"id"`
	CardHolderName string    `json:"card_holder_name"`
	CardNumber     string    `json:"card_number"`
	ExpirationDate string    `json:"expiration_date"`
	CreatedAt      time.Time `json:"created_at"`
}

// Хранимые данные пользователя - текстовое сообщение
type TextMessage struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// Хранимые данные пользователя - двоичное сообщение
type BinaryMessage struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	FileName  string    `json:"file_name"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
}
