package utils

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/golang-module/dongle"
)

func CheckAESKeyValid(input string) bool {
	keyBytes := []byte(input)
	keyLen := len(keyBytes)

	switch keyLen {
	case 16, 24, 32:
		return true
	default:
		return false
	}
}

func CheckAESIVValid(input string) bool {
	keyBytes := []byte(input)
	keyLen := len(keyBytes)

	switch keyLen {
	case 16:
		return true
	default:
		return false
	}
}

func GenAESRandomIV() string {
	mustLen := 16

	return RandomString(mustLen)
}

func AESEncrypt(key string, iv string, data string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			result = ""
			err = fmt.Errorf("aes encrypt error : %v", r)
			return
		}
	}()

	if !CheckAESKeyValid(key) {
		return "", errors.New("key长度错误")
	}
	if !CheckAESIVValid(key) {
		return "", errors.New("iv长度错误")
	}

	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)      // CBC、CFB、OFB、CTR、ECB
	cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
	cipher.SetKey(key)              // key must be 16, 24 or 32 bytes
	cipher.SetIV(iv)                // iv must be 16 bytes (ECB mode doesn't require setting iv)

	// Encrypt by aes from string and output raw string
	result = dongle.Encrypt.FromString(data).ByAes(cipher).ToBase64String()
	return
}

func AESDecrypt(key string, iv string, data string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			result = ""
			err = fmt.Errorf("aes decrypt error : %v", r)
			return
		}
	}()

	if !CheckAESKeyValid(key) {
		return "", errors.New("key长度错误")
	}
	if !CheckAESIVValid(key) {
		return "", errors.New("iv长度错误")
	}
	cipher := dongle.NewCipher()
	cipher.SetMode(dongle.CBC)      // CBC、CFB、OFB、CTR、ECB
	cipher.SetPadding(dongle.PKCS7) // No、Empty、Zero、PKCS5、PKCS7、AnsiX923、ISO97971
	cipher.SetKey(key)              // key must be 16, 24 or 32 bytes
	cipher.SetIV(iv)                // iv must be 16 bytes (ECB mode doesn't require setting iv)
	result = dongle.Decrypt.FromBase64String(data).ByAes(cipher).ToString()
	return
}
