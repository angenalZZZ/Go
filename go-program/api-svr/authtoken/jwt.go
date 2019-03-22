package authtoken

import (
	"angenalZZZ/go-program/api-config"
	"angenalZZZ/go-program/api-svr/cors"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

var (
	KID        string
	Audience   string
	alg        jwt.SigningMethod
	isEs, isRs bool
)

type PayloadClaims struct {
	jwt.StandardClaims
	Data string `json:"data"`
}

func init() {
	api_config.LoadCheck()

	KID = "0cb3c5b637d4bf3"
	Audience = "https://jwt.io, https://fpapi.com"
	switch api_config.JwtConf.JWT_algorithms {
	case "HS384":
		alg = jwt.SigningMethodHS384
	case "HS512":
		alg = jwt.SigningMethodHS512
	case "RS256":
		alg = jwt.SigningMethodRS256
	case "RS384":
		alg = jwt.SigningMethodRS384
	case "RS512":
		alg = jwt.SigningMethodRS512
	default:
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
		FError(&w, errors.New("账号信息认证失败"), true)
		return
	}

	// 签发token
	id := uuid.Must(uuid.NewV4()).String()
	now := time.Now()
	esAt := now.Unix()
	expAt := now.Add(time.Duration(api_config.JwtConf.JWT_LIFETIME) * time.Second).Unix()
	claims := PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "",       // 签发者
			Subject:   "",       // 面向的用户
			Audience:  Audience, // 接收 JWT 的一方(域名或应用名)
			ExpiresAt: expAt,    // 过期时间
			Id:        id,       // JWT 的唯一身份标识，主要用来作为一次性 token，从而避免重放攻击
			IssuedAt:  esAt,     // JWT 签发时间
			NotBefore: esAt,     // 什么时间之前，该 JWT 都是不可用的
		},
		Data: data,
	}
	//h := jwt.Header{KeyID: KID}
	key := []byte(api_config.JwtConf.JWT_SECRET)
	token := jwt.NewWithClaims(alg, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		FError(&w, err, true)
		return
	}
	FOk(&w, []byte(tokenString), true)
}

/**
curl -X POST http://localhost:8008/token/verify -H 'Content-Type: application/json' -d '{}' -H 'Authorization: bearer '
*/
// Verifying and validating a JWT http handler
func JwtVerifyValidateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(&w, r, []string{http.MethodPost}) {
		return
	}

	//log.Printf(" http.Request.Header: \n  %+v \n", r.Header)
	tokenString := r.Header.Get("authorization")
	if i := strings.Index(tokenString, " "); i > 0 {
		tokenString = tokenString[i+1:]
	}
	//log.Printf(" http.Request.Header.BearerToken: \n  %+v \n", tokenString)
	if len(tokenString) < 100 {
		FError(&w, jwt.NewValidationError("token format err!", jwt.ValidationErrorMalformed), true)
		return
	}

	// parsing token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the alg
		if isEs {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError(fmt.Sprintf("token alg(%v) err!", token.Header["alg"]), jwt.ValidationErrorSignatureInvalid)
			}
		}
		if isRs {
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
		FError(&w, err, true)
		return
	}

	// output token
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		FOk(&w, struct {
			header  map[string]interface{}
			payload jwt.Claims
		}{token.Header, token.Claims}, true)
	} else {
		FError(&w, jwt.NewValidationError("token err!", jwt.ValidationErrorNotValidYet), true)
	}
}

// response ok
func FOk(response *http.ResponseWriter, token interface{}, outputJson bool) {
	w := *response
	if outputJson == true {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		body := map[string]interface{}{"code": 0, "token": token, "msg": "success"}
		json.NewEncoder(w).Encode(body)
	} else {
		fmt.Fprint(w, token)
	}
}

// response error
func FError(response *http.ResponseWriter, err error, outputJson bool) {
	w := *response
	//set json response
	w.WriteHeader(http.StatusAccepted)
	if outputJson == true {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		body := map[string]interface{}{"code": 1, "token": "", "msg": fmt.Sprintf("%v", err)}
		json.NewEncoder(w).Encode(body)
	} else {
		fmt.Fprintf(w, "%v", err)
	}
}
