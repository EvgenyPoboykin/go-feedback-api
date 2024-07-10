package auth

import (
	"context"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		authorizationHeader := r.Header.Get("Authorization")
		if len(authorizationHeader) == 0 {
			permissionDenied(w, authHeaderError)

			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			permissionDenied(w, authFormatHeaderError)

			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeKey {
			permissionDenied(w, authUnSupportTypeError+authorizationType)

			return
		}

		accessToken := fields[1]
		userInfo := ParseToken(accessToken)
		if userInfo == nil {
			serviceUnavailable(w, notAuthorizedError)

			return
		}

		roleValue := containsRole(userInfo.RealmAccess.Roles)
		if roleValue == "" {
			serviceUnavailable(w, notAuthorizedError)

			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, role, roleValue)
		ctx = context.WithValue(ctx, clientId, userInfo.ClientId)
		ctx = context.WithValue(ctx, fullName, userInfo.FamilyName+" "+userInfo.GivenName+" "+userInfo.MiddleName)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
