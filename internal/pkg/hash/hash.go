package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func xorFunction(input, key string) string {
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ key[i%len(key)]
	}

	return string(output)
}

func generateTimestamp(duration string) int64 {
	if duration == "0" {
		return 0
	}
	d, _ := time.ParseDuration(duration)

	return time.Now().Add(d).Unix()
}

func generateNonce() string {
	rand.Seed(time.Now().UnixNano())
	nonce := make([]byte, 8)
	for i := range nonce {
		nonce[i] = byte(rand.Intn(256))
	}

	return base64.StdEncoding.EncodeToString(nonce)
}

func EncryptionURL(url, expStr, secretKey string) (string, error) {
	timestamp := generateTimestamp(expStr)
	nonce := generateNonce()
	combined := nonce + "|" + url + "|" + strconv.FormatInt(timestamp, 10)
	encrypted := xorFunction(combined, secretKey)

	return base64.StdEncoding.EncodeToString([]byte(encrypted)), nil
}

func DecryptionURL(hashURL, secretKey string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(hashURL)
	if err != nil {
		return "", err
	}

	decrypted := xorFunction(string(decoded), secretKey)
	parts := strings.Split(decrypted, "|")
	if len(parts) != 3 {
		return "", errors.New("invalid hash")
	}

	nonce := parts[0]
	url := parts[1]
	timestamp, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return "", err
	}

	fmt.Println("Nonce:", nonce) // Debug uchun nonce ni ko'rsatamiz

	if timestamp != 0 && time.Now().Unix() > timestamp {
		return "", errors.New("URL has expired")
	}

	return url, nil
}

func main() {
	url := "https://example.com/files/test-file-url.png"
	secretKey := "this.is.secret"
	duration := "5m"

	hashURL, err := EncryptionURL(url, duration, secretKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(hashURL)

	decrypted, err := DecryptionURL(hashURL, secretKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(decrypted)
}

/*
1. 1-messageda chat yoq bo'lsa avtomatik chat ochish
2. 2 xil chat, bir-bir va guruh
3. chatda file yoki message yubora olishi kerak
4. rasmiy kutubxonadan foydalanish, gorilla ishlatish
*/
