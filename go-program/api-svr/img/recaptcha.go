package img

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

// 参考：https://www.cnblogs.com/stulzq/p/10714417.html
const (
	// V2 人机验证
	recaptchaPublicKey  = "6LeOGZ8UAAAAAASOnKq1Lzvo6eNo7nky4a0nl5AJ"
	recaptchaPrivateKey = "6LeOGZ8UAAAAAGyxNdp2wxtvfr22EtZOaiaYDs1d"
	// V3
	//recaptchaPublicKey  = "6Lcw6p4UAAAAAA-_WY2tdq7LPVnB06dWsLccvPv5"
	//recaptchaPrivateKey = "6Lcw6p4UAAAAAKP-z8GbltwzGcOYYpBDnSqSkyQ-"
)

func init() {
	// http://localhost:8008/api/captcha/recaptcha
	http.HandleFunc("/api/captcha/recaptcha", RecaptchaGenerateHandler)
}

const (
	pageTop = `<!DOCTYPE HTML><html><head>
<style>.error{color:#ff0000;} .ack{color:#0000ff;}</style><title>Recaptcha Test</title></head>
<body><div style="width:100%"><div style="width: 50%;margin: 0 auto;">
<h3>Recaptcha Test</h3>
<p>This is a test form for the go-recaptcha package</p>`
	form = `<form action="" method="POST">
	    <script src="https://recaptcha.net/recaptcha/api.js"></script>
			<div class="g-recaptcha" data-sitekey="%s"></div>
    	<input type="submit" name="button" value="Ok">
</form>`
	pageBottom = `</div></div></body></html>`
	anError    = `<p class="error">%s</p>`
	anAck      = `<p class="ack">%s</p>`
)

func RecaptchaGenerateHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // Must be called before writing response
	fmt.Fprint(writer, pageTop)
	if err != nil {
		fmt.Fprintf(writer, fmt.Sprintf(anError, err))
	} else {
		_, buttonClicked := request.Form["button"]
		if buttonClicked {
			if recaptchaVerifyHandle(request) {
				fmt.Fprint(writer, fmt.Sprintf(anAck, "Recaptcha was correct!"))
			} else {
				fmt.Fprintf(writer, fmt.Sprintf(anError, "Recaptcha was incorrect; try again."))
			}
		}
	}
	fmt.Fprint(writer, fmt.Sprintf(form, recaptchaPublicKey))
	fmt.Fprint(writer, pageBottom)
}

func recaptchaVerifyHandle(request *http.Request) (result bool) {
	recaptchaResponse, responseFound := request.Form["g-recaptcha-response"]
	if responseFound {
		result, err := confirm("127.0.0.1", recaptchaResponse[0])
		if err != nil {
			log.Println("recaptcha server error", err)
		}
		return result
	}
	return false
}

/////////////////////////////////////////////////////////////////

type recaptchaResponse struct {
	Success     bool      `json:"success"`
	Score       float64   `json:"score"`
	Action      string    `json:"action"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname"`
	ErrorCodes  []string  `json:"error-codes"`
}

// check uses the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func check(remoteip, response string) (r recaptchaResponse, err error) {
	resp, err := http.PostForm("https://recaptcha.google.cn/recaptcha/api/siteverify", // https://www.google.com/recaptcha/api/siteverify
		url.Values{"secret": {recaptchaPrivateKey}, "remoteip": {remoteip}, "response": {response}})
	if err != nil {
		log.Printf("Post error: %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Read error: could not read body: %s\n", err)
		return
	}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Printf("Read error: got invalid JSON: %s\n %s\n", err, body)
		return
	}
	return
}

// Confirm is the public interface function.
// It calls check, which the client ip address, the challenge code from the reCaptcha form,
// and the client's response input to that challenge to determine whether or not
// the client answered the reCaptcha input question correctly.
// It returns a boolean value indicating whether or not the client answered correctly.
func confirm(remoteip, response string) (result bool, err error) {
	resp, err := check(remoteip, response)
	result = resp.Success
	return
}
