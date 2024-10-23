package models

//  Пользователь
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

// Хранимые данные пользователя - логин и пароль
type LoginPassword struct {
	Login     string `json:"login"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
}

// Хранимые данные пользователя - банковская карта
type BankCard struct {
	ID             int64  `json:"id"`
	CardHolderName string `json:"card_holder_name"`
	CardNumber     string `json:"card_number"`
	ExpirationDate string `json:"expiration_date"`
	CreatedAt      int64  `json:"created_at"`
}

// Хранимые данные пользователя - текстовое сообщение
type TextMessage struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"created_at"`
}

// Хранимые данные пользователя - двоичное сообщение
type BinaryMessage struct {
	ID        int64  `json:"id"`
	Binary    string `json:"binary"`
	CreatedAt int64  `json:"created_at"`
}
