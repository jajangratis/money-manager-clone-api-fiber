package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Config struct for holding configurations
func secEncrypt(key []byte, data interface{}, config Config) (string, error) {
	encoded, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, config.LoginIV)
	ciphertext := make([]byte, len(encoded))
	cfb.XORKeyStream(ciphertext, encoded)

	return fmt.Sprintf("%x", ciphertext), nil
}

func secDecrypt(key []byte, data string, config Config) (map[string]interface{}, error) {
	ciphertext, err := hex.DecodeString(data)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decipher := cipher.NewCFBDecrypter(block, config.LoginIV)
	decrypted := make([]byte, len(ciphertext))
	decipher.XORKeyStream(decrypted, ciphertext)

	var result map[string]interface{}
	err = json.Unmarshal(decrypted, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Encoded(input string) string {
	godotenv.Load()
	key := []byte(os.Getenv("KEY_LOGIN")) // Replace with your actual key
	data := map[string]interface{}{"example": input}

	config := Config{
		LoginIV: []byte(os.Getenv("LOGIN_IV")), // Replace with your actual IV
	}

	encrypted, err := secEncrypt(key, data, config)
	if err != nil {
		fmt.Println("Error encrypting:", err)
		PanicIfError(err)
	}

	fmt.Println("Encrypted:", encrypted)
	return encrypted
}

func Decoded(input string) interface{} {
	godotenv.Load()
	key := []byte(os.Getenv("KEY_LOGIN")) // Replace with your actual key

	config := Config{
		LoginIV: []byte(os.Getenv("LOGIN_IV")), // Replace with your actual IV
	}
	decrypted, err := secDecrypt(key, input, config)
	if err != nil {
		fmt.Println("Error decrypting:", err)
		PanicIfError(err)
	}

	//fmt.Println("Decrypted:", decrypted["example"])
	return decrypted["example"]
}
