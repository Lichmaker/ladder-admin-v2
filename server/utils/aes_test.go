package utils

import "testing"

func TestAESEncrypt(t *testing.T) {
	key := "1234567890qwerty"
	iv := "amG3DsI8yqDkHmRW"
	data := "qwert"

	encryptData, err := AESEncrypt(key, iv, data)
	if err != nil {
		t.Error(err)
	}
	if encryptData != "l/JqPmcoVO+EnIwpnS/z+g==" {
		t.Error("incorrect")
	}
}

func TestAESDecrypt(t *testing.T) {
	key := "1234567890qwerty"
	iv := "amG3DsI8yqDkHmRW"
	data := "l/JqPmcoVO+EnIwpnS/z+g=="

	rawData, err := AESDecrypt(key, iv, data)
	if err != nil {
		t.Error(err)
	}
	if rawData != "qwert" {
		t.Error("incorrect")
	}
}
