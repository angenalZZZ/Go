package authtoken

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	api_config "github.com/angenalZZZ/Go/go-program/api-config"
	"github.com/angenalZZZ/Go/go-program/api-svr/cors"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/xid"
)

/**
Jwt-Token数据
*/
var (
	// 签发者
	Issuer string
	// 面向的用户
	Subject string
	// 接收 JWT 的一方(域名或应用名)
	Audience string
	// the token alg
	tokenAlg             jwt.SigningMethod
	tokenIsHs, tokenIsRs bool
)

/**
Jwt-JSON数据签名
*/
var (
	// signing key and pub
	signKeyData, signPubData []byte
	// the signing key alg
	signAlg            jwt.SigningMethod
	signHasKeyAndPub   bool
	signIsHs, signIsRs bool
)

/**
Jwt-Claims数据结构
*/
type PayloadClaims struct {
	jwt.StandardClaims
	X string `json:"x"`
}

// 初始化参数
func init() {

	JwtConf := api_config.Config.Jwt
	//fmt.Printf(" %#v \n", JwtConf)

	Issuer = JwtConf.JWT_Issuer
	Subject = JwtConf.JWT_Subject
	Audience = JwtConf.JWT_Audience

	if JwtConf.JWT_Sign.HasKeyAndPub() {
		p := os.Getenv("GOPATH") + "/src/github.com/angenalZZZ/Go/go-program/"

		signKeyData, _ = api_config.LoadArgInput(p + JwtConf.JWT_Sign.Key)
		signPubData, _ = api_config.LoadArgInput(p + JwtConf.JWT_Sign.Pub)

		// get the signing key alg
		signAlg = jwt.GetSigningMethod(JwtConf.JWT_Sign.Alg)
		signHasKeyAndPub = len(signKeyData) > 1
		signIsHs = strings.Contains(strings.ToUpper(JwtConf.JWT_Sign.Key), "HS")
		signIsRs = strings.Contains(strings.ToUpper(JwtConf.JWT_Sign.Key), "RS")

		//fmt.Printf("  %v\n  %v\n  %v\n  %v\n  %v\n  %v\n", signKeyData, signPubData, signAlg, signHasKeyAndPub, signIsHs, signIsRs)
	}

	// get the token alg
	tokenAlg = jwt.GetSigningMethod(JwtConf.JWT_algorithms)
	tokenIsHs = strings.HasPrefix(JwtConf.JWT_algorithms, "HS")
	tokenIsRs = strings.HasPrefix(JwtConf.JWT_algorithms, "RS")

	//fmt.Printf("  %v\n  %v\n  %v\n", tokenAlg, tokenIsHs, tokenIsRs)

	// 账号信息认证：AUTH JWT
	http.HandleFunc("/token/jwt", JwtTokenGenerateHandler)
	http.HandleFunc("/token/jwt/verify", JwtVerifyValidateHandler)
	http.HandleFunc("/token/jwt/sign", JsonSignGenerateHandler)
	http.HandleFunc("/token/jwt/sign/verify", JsonSignValidateHandler)
}

/**
账号信息认证
*/
func accountValidate(r *http.Request) (data string, ok bool) {
	ok = strings.Contains(r.URL.Path, "token")
	data = r.URL.Query().Get("id")
	return
}

/**
Jwt-Token数据生成
curl -X GET http://localhost:8008/token/jwt?id=1553155268644
*/
func JwtTokenGenerateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(w, r, []string{http.MethodGet, http.MethodPost}) {
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
	esAt, expAt := now.Unix(), now.Add(time.Duration(api_config.Config.Jwt.JWT_LIFETIME)*time.Second).Unix()
	claims := PayloadClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    Issuer,   // 签发者
			Subject:   Subject,  // 面向的用户
			Audience:  Audience, // 接收 JWT 的一方(域名或应用名)
			ExpiresAt: expAt,    // 过期时间
			Id:        id,       // JWT 的唯一身份标识，主要用来作为一次性 token，从而避免重放攻击
			IssuedAt:  esAt,     // JWT 签发时间
			NotBefore: esAt,     // 什么时间之前，该 JWT 都是不可用的
		},
		X: data, // 扩展信息
	}
	key := []byte(api_config.Config.Jwt.JWT_SECRET)
	token := jwt.NewWithClaims(tokenAlg, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write([]byte(tokenString))
}

/**
Jwt-Token数据验证
curl -X POST http://localhost:8008/token/jwt/verify -H "Content-Type:application/json" -H "Authorization: Bearer {0}"
*/
func JwtVerifyValidateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(w, r, []string{http.MethodPost}) {
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
		// validate the tokenAlg
		if tokenIsHs {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError(fmt.Sprintf("token tokenAlg(%v) err!", token.Header["tokenAlg"]), jwt.ValidationErrorSignatureInvalid)
			}
		} else if tokenIsRs {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, jwt.NewValidationError(fmt.Sprintf("token tokenAlg(%v) err!", token.Header["tokenAlg"]), jwt.ValidationErrorSignatureInvalid)
			}
		}

		// secret is a []byte containing your secret, e.g. []byte("my_secret_key")
		secret := []byte(api_config.Config.Jwt.JWT_SECRET)
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

/**
Jwt-JSON数据签名生成
curl -X POST http://localhost:8008/token/jwt/sign -H "Content-Type:application/json" -d "{\"id\":\"1553155268644\"}"
*/
func JsonSignGenerateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(w, r, []string{http.MethodPost}) {
		return
	}
	if signHasKeyAndPub == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 加载签名 jwt key
	var err error
	var key interface{} = signKeyData[:]
	k, ok := key.([]byte)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//parse request parameters
	var claims jwt.MapClaims
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err = decoder.Decode(&claims); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// create a new token
	token := jwt.NewWithClaims(signAlg, claims)

	if signIsRs {
		if key, err = jwt.ParseRSAPrivateKeyFromPEM(k); err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	} else if signIsHs {
		if key, err = jwt.ParseECPrivateKeyFromPEM(k); err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	tokenString, e := token.SignedString(key)
	if e != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Write([]byte(tokenString))
}

/**
Jwt-JSON数据签名验证
curl -X POST http://localhost:8008/token/jwt/sign/verify -H "Content-Type:application/json" -H "Authorization: Bearer {0}"
*/
func JsonSignValidateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(w, r, []string{http.MethodPost}) {
		return
	}
	if signHasKeyAndPub == false {
		w.WriteHeader(http.StatusBadRequest)
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
		if signIsRs {
			return jwt.ParseRSAPublicKeyFromPEM(signPubData)
		} else if signIsHs {
			return jwt.ParseECPublicKeyFromPEM(signPubData)
		}
		return signKeyData, nil
	})

	// parsing the error
	if err != nil {
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
