package users

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Id       string `json:"id"`
	City     string `json:"city"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

/*
We need to implement hashing and salting
(so that even the encrypted passwords cannot be reverse-engineered easily through rainbow tables or brute force attacks)\
For the purpose of this project, simply encrypting the passwords using a cryptographic algorithm will suffice,
Whereas in practice we would use hashing as well as salting with more secure algorithms
*/
//secret-key
const passphrase string = "srikanth.balakrishna511@gmail.com"

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (user *User) EncryptPassword() {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	user.Password = string(gcm.Seal(nonce, nonce, []byte(user.Password), nil))
	fmt.Println("User encrypted password:" + user.Password)

}
