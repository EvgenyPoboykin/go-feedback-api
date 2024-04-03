package authmiddleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/eugenepoboykin/go-feedback-api/constant"
	"github.com/eugenepoboykin/go-feedback-api/helpers"
	"github.com/eugenepoboykin/go-feedback-api/utils"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtKey = "secret"

	authorizationHeaderKey = "Authorization"
	authorizationTypeKey   = "bearer"

	role     = "oauth.role"
	fullName = "oauth.fullName"
	clienyId = "oauth.clientId"
)

func containsRole(s []string, e string) string {
	for _, a := range s {
		if a == constant.Employee {
			return "employee"
		}

		if a == constant.Admin {
			return "admin"
		}
	}
	return ""
}

func getRole(userInfo UserToken) string {
	return containsRole(userInfo.RealmAccess.Roles, constant.Employee)
}

func getJwtKey() []byte {
	envJwt := utils.GetEnv("JWT_SECRET_KEY", jwtKey)

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
			helpers.ErrorResponse(w, http.StatusUnauthorized, constant.NOT_AUTHORIZED, constant.ResponseMessage_AuthHeaderError)

			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			helpers.ErrorResponse(w, http.StatusUnauthorized, constant.NOT_AUTHORIZED, constant.ResponseMessage_AuthFormatHeaderError)

			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeKey {
			helpers.ErrorResponse(w, http.StatusUnauthorized, constant.NOT_AUTHORIZED, constant.ResponseMessage_AuthUnsupportTypeError+authorizationType)

			return
		}

		accessToken := fields[1]
		userInfo := ParseToken(accessToken)
		if userInfo == nil {
			helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.NOT_AUTHORIZED, constant.ResponseMessage_NotAuthorizedError)

			return
		}

		roleValue := getRole(*userInfo)
		if roleValue == "" {
			helpers.ErrorResponse(w, http.StatusServiceUnavailable, constant.NOT_AUTHORIZED, constant.ResponseMessage_NotAuthorizedError)

			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, role, roleValue)
		ctx = context.WithValue(ctx, clienyId, userInfo.ClientId)
		ctx = context.WithValue(ctx, fullName, userInfo.FamilyName+" "+userInfo.GivenName+" "+userInfo.MiddleName)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
