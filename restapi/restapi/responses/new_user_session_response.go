package responses

import (
	"crypto/subtle"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//go:generate easytags $GOFILE json,xml

// custom github.com/dgrijalva/jwt-go@v3.2.0+incompatible/map_claims.go:10
type NewUserSession struct {
	Pubkey      string          `json:"pubkey" xml:"pubkey"`
	Uid         int64           `json:"uid" xml:"uid"`
	MapLock     map[string]bool `json:"map_lock" xml:"map_lock"`
	Token       string          `json:"token" xml:"token"`
	Duration    int64           `json:"duration" xml:"duration"`
	TimeExpired int64           `json:"time_expired" xml:"time_expired"`
	CreatedAt   int64           `json:"created_at" xml:"created_at"`
}

//// Compares the aud claim against cmp.
//// If required is false, this method will return true if the value matches or is unset
//func (m NewUserSession) VerifyAudience(cmp string, req bool) bool {
//	aud, _ := m["aud"].(string)
//	return verifyAud(aud, cmp, req)
//}

// Compares the exp claim against cmp.
// If required is false, this method will return true if the value matches or is unset
func (m NewUserSession) VerifyExpiresAt(cmp int64, req bool) bool {
	return verifyExp(m.TimeExpired, cmp, req)
}

// Compares the iat claim against cmp.
// If required is false, this method will return true if the value matches or is unset
//func (m NewUserSession) VerifyIssuedAt(cmp int64, req bool) bool {
//	switch iat := m["iat"].(type) {
//	case float64:
//		return verifyIat(int64(iat), cmp, req)
//	case json.Number:
//		v, _ := iat.Int64()
//		return verifyIat(v, cmp, req)
//	}
//	return req == false
//}

// Compares the iss claim against cmp.
// If required is false, this method will return true if the value matches or is unset
//func (m NewUserSession) VerifyIssuer(cmp string, req bool) bool {
//	iss, _ := m["iss"].(string)
//	return verifyIss(iss, cmp, req)
//}

// Compares the nbf claim against cmp.
// If required is false, this method will return true if the value matches or is unset
//func (m NewUserSession) VerifyNotBefore(cmp int64, req bool) bool {
//	switch nbf := m["nbf"].(type) {
//	case float64:
//		return verifyNbf(int64(nbf), cmp, req)
//	case json.Number:
//		v, _ := nbf.Int64()
//		return verifyNbf(v, cmp, req)
//	}
//	return req == false
//}

// Validates time based claims "exp, iat, nbf".
// There is no accounting for clock skew.
// As well, if any of the above claims are not in the token, it will still
// be considered a valid claim.
func (m NewUserSession) Valid() error {
	vErr := new(jwt.ValidationError)
	now := time.Now().UnixNano() / 1000

	if m.VerifyExpiresAt(now, false) == false {
		vErr.Inner = errors.New("Token is expired")
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	//if m.VerifyIssuedAt(now, false) == false {
	//	vErr.Inner = errors.New("Token used before issued")
	//	vErr.Errors |= jwt.ValidationErrorIssuedAt
	//}
	//
	//if m.VerifyNotBefore(now, false) == false {
	//	vErr.Inner = errors.New("Token is not valid yet")
	//	vErr.Errors |= jwt.ValidationErrorNotValidYet
	//}

	if vErr.Errors == 0 {
		return nil
	}

	return vErr
}

// ----- helpers

func verifyAud(aud string, cmp string, required bool) bool {
	if aud == "" {
		return !required
	}
	if subtle.ConstantTimeCompare([]byte(aud), []byte(cmp)) != 0 {
		return true
	} else {
		return false
	}
}

func verifyExp(exp int64, now int64, required bool) bool {
	if exp == 0 {
		return !required
	}
	return now <= exp
}

func verifyIat(iat int64, now int64, required bool) bool {
	if iat == 0 {
		return !required
	}
	return now >= iat
}

func verifyIss(iss string, cmp string, required bool) bool {
	if iss == "" {
		return !required
	}
	if subtle.ConstantTimeCompare([]byte(iss), []byte(cmp)) != 0 {
		return true
	} else {
		return false
	}
}

func verifyNbf(nbf int64, now int64, required bool) bool {
	if nbf == 0 {
		return !required
	}
	return now >= nbf
}
