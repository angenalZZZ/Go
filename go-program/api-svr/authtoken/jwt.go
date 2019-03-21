package authtoken

import (
	api_config "angenalZZZ/go-program/api-config"
	"angenalZZZ/go-program/api-svr/cors"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

var KID string
var Audience string
var hash *jwt.SigningMethod

type Payload struct {
	jwt.StandardClaims
	Uid string `json:"uid"`
}

func init() {
	api_config.LoadCheck()

	KID = "0cb3c5b637d4bf3"
	Audience = "https://jwt.io, https://fpapi.com"
	switch api_config.JwtConf.JWT_algorithms {
	case "HS384":
		hash = jwt.SigningMethodHS384
	case "HS512":
		hash = jwt.SigningMethodHS512
	default:
		hash = jwt.SigningMethodHS256
	}
}

// 账号信息认证
func accountValidate(r *http.Request) (uid string, ok bool) {
	ok = strings.Contains(r.URL.Path, "token")
	uid = r.URL.Query().Get("uid")
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
	uid, valid := accountValidate(r)
	if valid == false {
		FError(&w, errors.New("账号信息认证失败"), true)
	}

	// 签发token
	now := time.Now()
	claims := Payload{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "",                                                                           // 签发者
			Subject:   "",                                                                           // 面向的用户
			Audience:  Audience,                                                                     // 接收 JWT 的一方
			ExpiresAt: now.Add(time.Duration(api_config.JwtConf.JWT_LIFETIME) * time.Second).Unix(), // 过期时间
			NotBefore: now.Unix(),                                                                   // 什么时间之前，该 JWT 都是不可用的
			IssuedAt:  now.Unix(),                                                                   // JWT 签发时间
			Id:        uuid.Must(uuid.NewV4()).String(),                                             // JWT 的唯一身份标识，主要用来作为一次性 token，从而避免重放攻击
		},
		Uid: uid,
	}
	//h := jwt.Header{KeyID: KID}
	token := jwt.NewWithClaims(hash, claims)
	tokens, err := token.SignedString([]byte(api_config.JwtConf.JWT_SECRET))
	if err != nil {
		FError(&w, err, true)
	}
	FOk(&w, string(tokens), true)
}

/**
curl -X POST http://localhost:8008/token/verify -H 'Content-Type: application/json' -d '{}' -H 'Authorization: bearer '
*/
// Verifying and validating a JWT http handler
func JwtVerifyValidateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(&w, r, []string{http.MethodPost}) {
		return
	}

	bearerToken := r.Header.Get("Authorization")
	if strings.HasSuffix(bearerToken, "bearer ") {
		bearerToken = strings.TrimSuffix(bearerToken, "bearer ")
	}
	if len(bearerToken) < 100 || len(strings.Split(bearerToken, ".")) != 3 {
		FError(&w, errors.New("参数 Authorization 验证失败"), true)
	}

	// 跟踪请求
	fmt.Printf("Verifying token: %+v\n", bearerToken)

	now := time.Now()
	hs256 := jwt.NewHMAC(hash, []byte(api_config.JwtConf.JWT_SECRET))
	token := []byte(bearerToken)
	raw, err := jwt.Parse(token)
	if err != nil {
		FError(&w, err, true)
	}

	// 跟踪请求
	fmt.Printf("Verifying raw: %+v\n", raw)

	if err = raw.Verify(hs256); err != nil {
		FError(&w, err, true)
	}
	var (
		h jwt.Header
		p Payload
	)
	if h, err = raw.Decode(&p); err != nil {
		FError(&w, err, true)
	}
	if h.KeyID != KID {
		FError(&w, errors.New("参数 KeyID 验证失败"), true)
	}

	iatValidator := jwt.IssuedAtValidator(now)
	expValidator := jwt.ExpirationTimeValidator(now, true)
	audValidator := jwt.AudienceValidator(Audience)
	if err := p.Validate(iatValidator, expValidator, audValidator); err != nil {
		switch err {
		case jwt.ErrIatValidation:
			FError(&w, errors.New("参数 iat 验证失败"), true)
		case jwt.ErrExpValidation:
			FError(&w, errors.New("参数 exp 验证失败"), true)
		case jwt.ErrAudValidation:
			FError(&w, errors.New("参数 aud 验证失败"), true)
		default:
			FError(&w, err, true)
		}
		return
	}
	FOk(&w, bearerToken, true)
}

// response ok
func FOk(response *http.ResponseWriter, token string, outputJson bool) {
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
