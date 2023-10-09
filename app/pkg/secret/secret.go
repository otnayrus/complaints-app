package secret

import (
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (hash string, err error) {
	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}
	hash = string(hashedPassword)

	return
}

func MatchPassword(password string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errorwrapper.WrapErr(errorwrapper.ErrBadRequest, "password is incorrect")
		}
		return errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}
	return nil
}

func IsPasswordValid(password string, hash string) error {
	return MatchPassword(password, hash)
}
