package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/KznRkjp/go-keeper.git/internal/middleware/mlogger"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// тут непонятно
type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

// время действия токена
const TokenExp = time.Hour * 10

// константа используется для генерации - в идеале надо ее брать из env
const SecretKey = "supersecretkey"

// BuildJWTString - генерация JWT токена
func BuildJWTString(id int) (string, error) {
	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExp)),
		},
		// собственное утверждение
		UserID: id,
	})

	// создаём строку токена
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}

// получение UserID из токена
func GetUserID(tokenString string) (int, error) {
	// fmt.Println("****** starting jwt check")
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims,
		func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				// fmt.Println("Тут что то не так")
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
	if err != nil {
		return -2, err
	}

	if !token.Valid {
		mlogger.Info("Token is not valid")
		return -1, err
	}

	mlogger.Info("Token is valid")
	mlogger.Info(fmt.Sprintf("UserID: %d", claims.UserID))
	return claims.UserID, err
}

// HashPassword generates a bcrypt hash for the given password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword verifies if the given password matches the stored hash.
func VerifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func EncryptData(password string, message string) ([]byte, error) {
	src := []byte(message)
	key := sha256.Sum256([]byte(password))
	// NewCipher создает и возвращает новый cipher.Block.
	// Ключевым аргументом должен быть ключ AES, 16, 24 или 32 байта
	// для выбора AES-128, AES-192 или AES-256.
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}
	// NewGCM возвращает заданный 128-битный блочный шифр
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		mlogger.Info(err.Error())
		return nil, err
	}
	// создаём вектор инициализации
	nonce := key[len(key)-aesgcm.NonceSize():]

	dst := aesgcm.Seal(nil, nonce, src, nil) // зашифровываем
	mlogger.Info("encrypted: " + string(dst))
	return dst, nil
}

func DecryptData(password string, message []byte) (string, error) {
	src := message
	mlogger.Info("encrypted: " + string(src))
	key := sha256.Sum256([]byte(password))
	// NewCipher создает и возвращает новый cipher.Block.
	// Ключевым аргументом должен быть ключ AES, 16, 24 или 32 байта
	// для выбора AES-128, AES-192 или AES-256.
	aesblock, err := aes.NewCipher(key[:])
	if err != nil {
		mlogger.Info(err.Error())
		return "", err
	}
	// NewGCM возвращает заданный 128-битный блочный шифр
	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		mlogger.Info(err.Error())
		return "", err
	}
	// создаём вектор инициализации
	nonce := key[len(key)-aesgcm.NonceSize():]

	src2, err := aesgcm.Open(nil, nonce, src, nil) // расшифровываем
	if err != nil {
		mlogger.Info(err.Error())
		return "", err
	}
	mlogger.Info("decrypted: " + string(src2))
	return string(src2), nil
}
