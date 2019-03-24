package api_config

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	//"github.com/chai2010/winsvc"
	"github.com/gobuffalo/envy"
	"log"
	"os"
	"strconv"
)

var JwtConf *JwtConfig

// 认证方式 Jwt Config: 复合类型
type JwtConfig struct {
	AUTH_JWT, JWT_algorithms, JWT_SECRET  string
	JWT_LIFETIME                          int
	JWT_Issuer, JWT_Subject, JWT_Audience string
	JWT_Sign                              JwtSign
}
type JwtSign struct {
	Key, Pub, Alg string
}

func (o JwtSign) HasKeyAndPub() bool {
	return o.Key != "" && o.Pub != "" && o.Alg != ""
}

// 加载配置文件
func init() {

	// *** 文件 .env 编码 必须是 UTF-8 +换行LF ***
	//AppPath, _ = winsvc.GetAppPath()

	// 检查环境，判断加载的配置文件
	f, v := ".env.prod", os.Getenv("GO_ENV")
	if v == "development" {
		f = ".env"
	}
	//log.Printf("配置文件: 解析 %s\n", f)

	// 配置文件错误时，直接退出应用
	if e := envy.Load(f); e != nil {
		panic(e)
	} else if v == "development" {
		//for _, v := range envy.Environ() {
		//	println(v)
		//}
	}

}

// 加载输入参数|文件路径
func LoadArgInput(s string) ([]byte, error) {
	if s == "" {
		return nil, fmt.Errorf("no input")
	} else if s == "+" {
		return []byte("{}"), nil
	}
	var r io.Reader
	if s == "-" {
		r = os.Stdin
	} else {
		if f, e := os.Open(s); e != nil {
			return nil, e
		} else {
			defer f.Close() // end
			r = f
		}
	}
	return ioutil.ReadAll(r)
}

// "+" > {}
func JsonInput(s string) (v interface{}, e error) {
	var data []byte
	if data, e = LoadArgInput(s); e == nil {
		e = json.Unmarshal(data, v)
	}
	return
}

func JsonOutput(v interface{}, indent bool) (s []byte, e error) {
	if indent == false {
		s, e = json.MarshalIndent(v, "", "    ")
	} else {
		s, e = json.Marshal(v)
	}
	return
}

// 加载配置文件并检查配置项
func LoadCheck() {
	if JwtConf == nil {
		println()

		// 检查配置项目
		Check("AUTH_JWT")
		Check("JWT_algorithms")
		Check("JWT_SECRET")
		Check("JWT_LIFETIME")
		Check("JWT_Issuer")
		Check("JWT_Subject")
		Check("JWT_Audience")
		Check("JWT_Sign")

		lifetime, _ := strconv.Atoi(os.Getenv("JWT_LIFETIME"))

		jwtSign := JwtSign{}
		s := strings.Split(os.Getenv("JWT_Sign"), ",")
		if len(s) == 3 {
			jwtSign = JwtSign{s[0], s[1], s[2]}
		}

		JwtConf = &JwtConfig{
			AUTH_JWT:       os.Getenv("AUTH_JWT"),
			JWT_algorithms: os.Getenv("JWT_algorithms"),
			JWT_SECRET:     os.Getenv("JWT_SECRET"),
			JWT_LIFETIME:   lifetime,
			JWT_Issuer:     os.Getenv("JWT_Issuer"),
			JWT_Subject:    os.Getenv("JWT_Subject"),
			JWT_Audience:   os.Getenv("JWT_Audience"),
			JWT_Sign:       jwtSign,
		}

		log.Printf("加载配置文件并检查配置项: OK\n")
		//log.Printf("%#v \n", JwtConf)
	}
}

// 检查配置项 Must Checked
func Check(key string) {
	if _, e := envy.MustGet(key); e != nil {
		panic(e)
	} else {
		//log.Printf("配置文件: %s = %s \n", key, val)
	}
}
