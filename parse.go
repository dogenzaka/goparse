package goparse

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	me = "/users/me"
)

func GetMe(sessionToken string) (u User, err error) {

	if sessionToken == "" {
		err = errors.New("parse session token is empty")
		return
	}

	req := gorequest.New().Get(config.Url + me)
	initParseHeader(req)
	setParseSessionToken(req, sessionToken)

	resp, body, errs := req.Timeout(config.TimeOut).End()
	if errs != nil {
		err = errors.New(fmt.Sprintf("%v", errs))
		return
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = errors.New(resp.Status)
		return
	}

	err = json.NewDecoder(strings.NewReader(body)).Decode(&u)

	return
}

func initParseHeader(req *gorequest.SuperAgent) {
	req.Set("X-Parse-Application-Id", config.ApplicationId)
	req.Set("X-Parse-REST-API-Key", config.RESTAPIKey)
}

func setParseSessionToken(req *gorequest.SuperAgent, token string) {
	req.Set("X-Parse-Session-Token", token)
}
