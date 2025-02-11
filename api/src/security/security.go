package security

import "golang.org/x/crypto/bcrypt"

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerificarSenha(plainTextPwd, cipherTextPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(cipherTextPwd), []byte(plainTextPwd))
}
