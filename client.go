package synology

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/golang/glog"
)

// A Client is a Synology API Client.
type Client struct {
	Host       string
	Username   string
	Password   string
	httpClient *http.Client
}

func (c Client) get(path string, params map[string]string, knownErrors map[int]string) (interface{}, error) {
	if c.httpClient == nil {
		cookieJar, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
		c.httpClient = &http.Client{
			Jar: cookieJar,
		}
		_, err = c.login()
		if err != nil {
			return nil, err
		}
	}

	url := url.URL{
		Scheme: "http",
		Host:   c.Host,
		Path:   path,
	}
	query := url.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	url.RawQuery = strings.Replace(query.Encode(), "+", "%20", -1)
	glog.Infof("Request: %s", url.String())

	resp, err := c.httpClient.Get(url.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	glog.Infof("Response: %s", string(body))

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	err = errFromData(data, knownErrors)
	if err != nil {
		return data, err
	}

	return data.(map[string]interface{})["data"], nil
}

func (c Client) login() (interface{}, error) {
	return c.get(
		"/webapi/auth.cgi",
		map[string]string{
			"api":     "SYNO.API.Auth",
			"version": "2",
			"method":  "login",
			"account": c.Username,
			"passwd":  c.Password,
			"session": "CLI",
			"format":  "cookie",
		},
		map[int]string{
			400: "No such account or incorrect password",
			401: "Account disabled",
			402: "Permission denied",
			403: "2-step verification code required",
			404: "Failed to authenticate 2-step verification code",
		},
	)
}
