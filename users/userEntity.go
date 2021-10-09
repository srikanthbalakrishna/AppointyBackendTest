package users

import (
	"crypto/md5"
	"encoding/hex"
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
For the purpose of this project, simply hashing the passwords using a cryptographic algorithm will suffice,
Whereas in practice we would use hashing as well as salting with more secure algorithms
*/
func (user *User) EncryptPassword() {

	algorithm := md5.New()
	algorithm.Write([]byte(user.Password))
	user.Password = hex.EncodeToString(algorithm.Sum(nil))
}
