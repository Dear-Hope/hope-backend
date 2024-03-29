package helpers

import (
	"HOPE-backend/internal/constant"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func GetStartOfDayUTCFromDateWithOffset(date string, offset int) time.Time {
	t, _ := time.Parse(constant.FormatDate, date)
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, getTimezoneFromOffset(offset)).UTC()
}

func GetEndOfDayUTCFromDateWithOffset(date string, offset int) time.Time {
	t, _ := time.Parse(constant.FormatDate, date)
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, getTimezoneFromOffset(offset)).UTC()
}

func getTimezoneFromOffset(offset int) *time.Location {
	return time.FixedZone("", offset*constant.HourInSec)
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const secret = "this-is-32-secret-code-for-hope!"

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decode(s string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	return data, err
}

func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	cipherText, err := Decode(text)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func EncryptPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error encrypting password: %v", err)
	}

	return string(hash), nil
}

func TimestampToStringFormat(t time.Time, format string) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(format)
}
