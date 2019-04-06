package img

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/angenalZZZ/Go/go-program/api-svr/jsonp"

	"github.com/angenalZZZ/Go/go-program/api-svr/cors"
	"github.com/mojocn/base64Captcha"
)

/**
Build and Run the Demo: nginx config
server {
        listen 80;
        server_name captcha.mojotv.cn;
        charset utf-8;

        location / {
            try_files /_not_exists_ @backend;
        }
        location @backend {
           proxy_set_header X-Forwarded-For $remote_addr;
           pro=xy_set_header Host $http_host;
           proxy_pass http://127.0.0.1:8008;
        }
        access_log  /home/wwwlogs/captcha.mojotv.cn.log;
}
*/

// HTTP请求上下文
var ctx = &Base64Captcha{}

type Base64Captcha struct{}

// HTTP请求内容
var defaultConfig *ConfigJsonBody

type ConfigJsonBody struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

// 初始化HTTP请求处理：验证码
func init() {
	defaultConfig = &ConfigJsonBody{
		ConfigAudio: base64Captcha.ConfigAudio{CaptchaLen: 4, Language: "zh"},
		ConfigCharacter: base64Captcha.ConfigCharacter{
			Height: 40,
			Width:  120,
			//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合
			Mode:               base64Captcha.CaptchaModeArithmetic,
			ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
			ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
			IsUseSimpleFont:    true,
			IsShowHollowLine:   false,
			IsShowNoiseDot:     false,
			IsShowNoiseText:    false,
			IsShowSlimeLine:    false,
			IsShowSineLine:     false,
			CaptchaLen:         6,
		},
		ConfigDigit: base64Captcha.ConfigDigit{
			Height:     35,
			Width:      70,
			CaptchaLen: 4,
			MaxSkew:    0.8,
			DotCount:   60,
		},
	}

	http.HandleFunc("/api/captcha/get", ctx.CaptchaGenerateHandler)
	http.HandleFunc("/api/captcha/verify", ctx.CaptchaVerifyHandle)
}

// 生成新的验证码
func (ctx *Base64Captcha) CaptchaGenerateHandler(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(w, r, []string{http.MethodGet, http.MethodPost}) {
		return
	}

	// 处理请求参数：post-body请求完整内容、get-使用默认配置、标识id
	var cfg interface{}
	var body *ConfigJsonBody
	id := r.URL.Query().Get("id")
	if id == "" {
		id = r.URL.Query().Get("lastCode")
	}
	if r.Method == http.MethodPost {
		defer r.Body.Close()
		if e := json.NewDecoder(r.Body).Decode(body); e != nil {
			jsonp.Error(e).Error(w, r)
			return
		}
		switch body.CaptchaType {
		case "audio":
			cfg = body.ConfigAudio
		case "character":
			cfg = body.ConfigCharacter
		case "digit":
			cfg = body.ConfigDigit
		default:
			jsonp.Error(errors.New("参数异常")).Error(w, r)
			return
		}
		id = body.Id
	} else {
		captchaType := r.URL.Query().Get("captchaType")
		switch captchaType {
		case "audio":
			cfg = defaultConfig.ConfigAudio
		case "character":
			cfg = defaultConfig.ConfigCharacter
		case "digit":
			cfg = defaultConfig.ConfigDigit
		default:
			cfg = defaultConfig.ConfigDigit
		}
	}

	//create base64 encoding captcha
	newId, instance := base64Captcha.GenerateCaptcha(id, cfg)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(instance)

	//or you can just write the captcha content to the httpResponseWriter.
	//before you put the captchaId into the response COOKIE.
	//instance.WriteTo(w)

	//set response
	jsonp.Success(jsonp.Data{"data": base64blob, "id": newId}).OK(w, r)
}

// 验证验证码的正确性
func (ctx *Base64Captcha) CaptchaVerifyHandle(w http.ResponseWriter, r *http.Request) {
	if cors.Cors(w, r, []string{http.MethodPost}) {
		return
	}

	//parse request parameters
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var postParameters ConfigJsonBody
	body := map[string]interface{}{"code": 1} // response error
	err := decoder.Decode(&postParameters)
	if err == nil {
		id, verifyValue := postParameters.Id, postParameters.VerifyValue
		if id == "" {
			id = r.URL.Query().Get("id")
			if id == "" {
				id = r.URL.Query().Get("lastCode")
			}
		}
		if id != "" || verifyValue != "" {
			//verify the captcha
			verifyResult := base64Captcha.VerifyCaptcha(id, verifyValue)

			//set response
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			if verifyResult {
				body = map[string]interface{}{"code": 0} // response ok
			}
		}
	}
	json.NewEncoder(w).Encode(body)
}
