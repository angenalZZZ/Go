package authtoken

import (
	"angenalZZZ/go-program/api-config"
	"angenalZZZ/go-program/api-svr/cors"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/xid"
)

var (
	Audience   string
	alg        jwt.SigningMethod
	isHs, isRs bool
)

type PayloadClaims struct {
	jwt.StandardClaims
	X string `json:"x"`
}

func init() {
	api_config.LoadCheck()

	Audience = "fpapi.com"

	switch api_config.JwtConf.JWT_algorithms {
	case "HS384":
		isHs = true
		alg = jwt.SigningMethodHS384
	case "HS512":
		isHs = true
		alg = jwt.SigningMethodHS512
	case "RS256":
		isRs = true
		alg = jwt.SigningMethodRS256
	case "RS384":
		isRs = true
		alg = jwt.SigningMethodRS384
	case "RS512":
		isRs = true
		alg = jwt.SigningMethodRS512
	default:
		isHs = true
		alg = jwt.SigningMethodHS256
	}
}

// 账号信息认证
func accountValidate(r *http.Request) (data string, ok bool) {
	ok = strings.Contains(r.URL.Path, "token")
	data = r.URL.Query().Get("id")
	return
}

/**
curl -X GET http://localhost:8008/token?id=1553155268644
*/
// Signing a JWT with public claims http handler
func JwtTokenGenerateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(&w, r, []string{http.MethodGet, http.MethodPost}) {
		return
	}

	// 账号信息认证
	data, valid := accountValidate(r)
	if valid == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 签发token
	id := xid.New().String()
	now := time.Now()
	esAt, expAt := now.Unix(), now.Add(time.Duration(api_config.JwtConf.JWT_LIFETIME)*time.Second).Unix()
	claims := PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "api-jwt",    // 签发者
			Subject:   "auth-token", // 面向的用户
			Audience:  Audience,     // 接收 JWT 的一方(域名或应用名)
			ExpiresAt: expAt,        // 过期时间
			Id:        id,           // JWT 的唯一身份标识，主要用来作为一次性 token，从而避免重放攻击
			IssuedAt:  esAt,         // JWT 签发时间
			NotBefore: esAt,         // 什么时间之前，该 JWT 都是不可用的
		},
		X: data, // 扩展信息
	}
	key := []byte(api_config.JwtConf.JWT_SECRET)
	token := jwt.NewWithClaims(alg, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(tokenString))
}

/**
curl -X POST http://localhost:8008/token/verify -H 'Content-Type: application/json' -H 'Authorization: bearer '
*/
// Verifying and validating a JWT http handler
func JwtVerifyValidateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(&w, r, []string{http.MethodPost}) {
		return
	}

	//log.Printf(" http.Request.Header: \n  %+v \n", r.Header)
	tokenString := r.Header.Get("authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//tokenString = string(regexp.MustCompile(`\s*$`).ReplaceAll([]byte(tokenString), []byte{}))
	if i := strings.Index(tokenString, " "); i > 0 {
		tokenString = tokenString[i+1:]
	}
	//log.Printf(" http.Request.Header.BearerToken: \n  %+v \n", tokenString)
	if len(tokenString) < 100 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// parsing token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if isHs {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError(fmt.Sprintf("token alg(%v) err!", token.Header["alg"]), jwt.ValidationErrorSignatureInvalid)
			}
		} else if isRs {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, jwt.NewValidationError(fmt.Sprintf("token alg(%v) err!", token.Header["alg"]), jwt.ValidationErrorSignatureInvalid)
			}
		}

		// secret is a []byte containing your secret, e.g. []byte("my_secret_key")
		secret := []byte(api_config.JwtConf.JWT_SECRET)
		return secret, nil
	})

	// parsing the error types
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
			} else {
				// Couldn't handle this token
			}
		}
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// output token
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		body := struct {
			Header map[string]interface{} // The first segment of the token
			Claims jwt.Claims             // The second segment of the token
			Valid  bool
		}{token.Header, token.Claims, token.Valid}
		json.NewEncoder(w).Encode(body)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
