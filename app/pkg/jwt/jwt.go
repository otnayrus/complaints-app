package jwt

import (
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/otnayrus/sb-rest/app/pkg/errorwrapper"
)

// GenerateJWTStringWithClaims generates a JWT token with algorithm RS256.
func GenerateJWTStringWithClaims(claims map[string]interface{}) (string, error) {
	// set expiration time H+3
	claims["exp"] = jwt.TimeFunc().AddDate(0, 0, 3).Unix()

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))

	// sign token
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return "", errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	return tokenString, nil
}

// ParseJWTStringWithClaims parses a JWT token with algorithm RS256.
func ParseJWTStringWithClaims(tokenString string) (map[string]interface{}, error) {
	// separate bearer string
	auth := strings.Split(tokenString, " ")
	if len(auth) != 2 && strings.ToLower(auth[0]) != "bearer" {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrForbidden, "Invalid token format")
	}

	tokenString = auth[1]

	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	if err != nil {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, err.Error())
	}

	// validate token
	if !token.Valid {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrForbidden, "Invalid token")
	}

	// get claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrInternalServer, "Failed to parse claims")
	}

	// check expiration time
	if !claims.VerifyExpiresAt(jwt.TimeFunc().Unix(), true) {
		return nil, errorwrapper.WrapErr(errorwrapper.ErrForbidden, "Token expired")
	}

	return claims, nil
}
