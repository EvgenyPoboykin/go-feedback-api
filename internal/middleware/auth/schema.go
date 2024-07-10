package auth

import "github.com/golang-jwt/jwt/v5"

type Roles struct {
	Roles []string `json:"roles"`
}

type UserToken struct {
	ClientId    string `json:"clientId"`
	Name        string `json:"name"`
	KLogin      string `json:"k_login"`
	GivenName   string `json:"given_name"`
	MiddleName  string `json:"middle_name"`
	FamilyName  string `json:"family_name"`
	RealmAccess Roles  `json:"realm_access"`

	jwt.RegisteredClaims
}
