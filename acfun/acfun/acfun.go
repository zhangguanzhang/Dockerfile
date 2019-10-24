package acfun

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	domain   = "https://www.acfun.cn"
	LoginUrl = "https://id.app.acfun.cn/rest/web/login/signin"
	MaxTags  = 4
)

type Acfun struct {
	commentCount   uint
	ArticlePool    chan uint
	client         *resty.Client
	Username       string
	CertifiedValue string
	userId         int64
}

type baseInfo struct {
	Result   int64  `json:"result"` //接口返回json存在result:100000情况,不能为int8
	ErrorMsg string `json:"error_msg"`
}

type logInInfo struct {
	baseInfo
	UserID   int64  `json:"userId"`
	Img      string `json:"img"`
	Username string `json:"username"`
}

func NewAcfun(username, password string, debug ...string) (*Acfun, error) {
	var data = logInInfo{}
	ac := &Acfun{
		Username:    username,
		ArticlePool: make(chan uint, 100),
		client:      resty.New(),
	}
	if len(debug) != 0 && debug[0] != "false" && debug[0] != "" {
		ac.client.SetDebug(true)
	}
	resp, err := ac.R().SetFormData(map[string]string{
		"username": username,
		"password": password,
		"key":      "",
		"captcha":  "",
	}).Post(LoginUrl)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		return nil, fmt.Errorf("%s, %s", err, resp)
	}

	if data.Result != 0 {
		return nil, errors.New(data.ErrorMsg)
	}

	ac.Username = data.Username
	ac.userId = data.UserID
	rand.Seed(time.Now().UnixNano())
	ac.CertifiedValue = base64.StdEncoding.EncodeToString([]byte(strconv.FormatInt(rand.Int63(), 32)))
	ac.client.SetCookie(
		&http.Cookie{
			Name:   CertifiedCookieName,
			Value:  ac.CertifiedValue,
			Path:   "/",
			Domain: strings.Split(domain, "/")[1],
		})

	return ac, nil
}

func (ac *Acfun) R() *resty.Request {
	return ac.client.R().
		SetHeaders(map[string]string{
			"Accept":          "*/*",
			"Accept-encoding": "gzip, deflate, br",
			"Accept-language": "zh-CN,zh;q=0.9",
			"Authority":       "www.acfun.cn",
			"Referer":         "https://www.acfun.cn/member/",
			"Origin":          "https://www.acfun.cn",
			"Sec-Fetch-Mode":  "cors",
			"User-Agent":      "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.103 Safari/537.36",
		})
}

