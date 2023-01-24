package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"fmt"

)

func Encrypt(stringToEncrypt string) (encryptedString string) {
	// convert key to bytes
	// key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	// block, err := aes.NewCipher(key)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	// iv := ciphertext[:aes.BlockSize]
	// if _, err := io.ReadFull(rand.Reader, iv); err != nil {
	// 	panic(err)
	// }

	// stream := cipher.NewCFBEncrypter(iv)
	// stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(plaintext)

	// hasher := md5.New()
	// hasher.Write([]byte(key))
	// return hex.EncodeToString(hasher.Sum(nil))
}

func Decrypt(keyString string, stringToDecrypt string) string {
	key, _ := hex.DecodeString(keyString)
	ciphertext, _ := base64.URLEncoding.DecodeString(stringToDecrypt)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

// func Encrypt(plaintext string) (string, error) {
// 	h := sha256.New()
// 	if _, err := io.WriteString(h, plaintext); err != nil {
// 		return "", err
// 	}
// 	r := h.Sum(nil)
// 	return hex.EncodeToString(r), nil
// }
