package mw

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/eugenepoboykin/go-feedback-api/internal/env"
	"github.com/eugenepoboykin/go-feedback-api/internal/lib/response"

	"github.com/golang-jwt/jwt/v5"
)

const (
	authorizationHeaderKey = "Authorization"
	authorizationTypeKey   = "bearer"

	role     = "oauth.role"
	fullName = "oauth.fullName"
	clienyId = "oauth.clientId"

	NOT_AUTHORIZED = "NOT_AUTHORIZED"

	ResponseMessage_AuthHeaderError        = "Authorization header is not provided"
	ResponseMessage_AuthFormatHeaderError  = "Authorization header is not provided"
	ResponseMessage_AuthUnsupportTypeError = "Unsupported authorization type "
	ResponseMessage_NotAuthorizedError     = "You not authorized!"
)

func containsRole(s []string) string {
	for _, a := range s {
		if a == env.Environment.EmployeeRole {
			return "employee"
		}

		if a == env.Environment.AdminRole {
			return "admin"
		}
	}
	return ""
}

func getRole(userInfo UserToken) string {
	return containsRole(userInfo.RealmAccess.Roles)
}

func getJwtKey() []byte {
	envJwt, _ := os.LookupEnv("JWT_SECRET_KEY")

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

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authorizationHeader := r.Header.Get("Authorization")
		if len(authorizationHeader) == 0 {
			response.ErrorResponse(w, http.StatusUnauthorized, NOT_AUTHORIZED, ResponseMessage_AuthHeaderError)

			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			response.ErrorResponse(w, http.StatusUnauthorized, NOT_AUTHORIZED, ResponseMessage_AuthFormatHeaderError)

			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeKey {
			response.ErrorResponse(w, http.StatusUnauthorized, NOT_AUTHORIZED, ResponseMessage_AuthUnsupportTypeError+authorizationType)

			return
		}

		accessToken := fields[1]
		userInfo := ParseToken(accessToken)
		if userInfo == nil {
			response.ErrorResponse(w, http.StatusServiceUnavailable, NOT_AUTHORIZED, ResponseMessage_NotAuthorizedError)

			return
		}

		roleValue := getRole(*userInfo)
		if roleValue == "" {
			response.ErrorResponse(w, http.StatusServiceUnavailable, NOT_AUTHORIZED, ResponseMessage_NotAuthorizedError)

			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, role, roleValue)
		ctx = context.WithValue(ctx, clienyId, userInfo.ClientId)
		ctx = context.WithValue(ctx, fullName, userInfo.FamilyName+" "+userInfo.GivenName+" "+userInfo.MiddleName)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
