package auth

import (
	"fmt"
	"github.com/eugenepoboykin/go-feedback-api/internal/domain/env"
	"github.com/eugenepoboykin/go-feedback-api/internal/errors"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

const (
	authorizationTypeKey = "bearer"

	role     = "oauth.role"
	fullName = "oauth.fullName"
	clientId = "oauth.clientId"

	notAuthorized = "NOT_AUTHORIZED"

	authHeaderError        = "Authorization header is not provided"
	authFormatHeaderError  = "Authorization header is not provided"
	authUnSupportTypeError = "Unsupported authorization type "
	notAuthorizedError     = "You not authorized!"
)

func containsRole(s []string) string {
	for _, a := range s {
		if a == env.Environment.EmployeeRole {
			return env.Environment.EmployeeRole
		}

		if a == env.Environment.AdminRole {
			return env.Environment.AdminRole
		}
	}
	return ""
}

func getJwtKey() []byte {
	envJwt := env.Environment.Secret

	return []byte(envJwt)
}

func ParseToken(t string) *UserToken {

	token, err := jwt.ParseWithClaims(t, &UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return getJwtKey(), nil
	})

	if err != nil {
		fmt.Print(err)
		return nil
	}

	claims := token.Claims.(*UserToken)

	return claims

}

func permissionDenied(w http.ResponseWriter, text string) {
	response.ErrorResponse(w, http.StatusUnauthorized, *errors.Error(notAuthorized, text))
}

func serviceUnavailable(w http.ResponseWriter, text string) {
	response.ErrorResponse(w, http.StatusServiceUnavailable, *errors.Error(notAuthorized, text))
}
