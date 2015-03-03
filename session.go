package goparse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	headerAppID        = "X-Parse-Application-Id" // Parse Application ID
	headerMasterKey    = "X-Parse-Master-Key"     // Parse Master Key
	headerAPIKey       = "X-Parse-REST-API-Key"   // Parse REST API Key
	headerSessionToken = "X-Parse-Session-Token"  // Parse Session Token

	pathMe = "/users/me"
)

// ParseSession is the client which has SessionToken as user authentication.
type ParseSession struct {
	client       *ParseClient
	SessionToken string
}

// Create a request which is set headers for Parse API
func (s *ParseSession) initRequest(req *gorequest.SuperAgent) {
	req.
		Set(headerAppID, s.client.ApplicationID).
		Set(headerAPIKey, s.client.RESTAPIKey).
		Timeout(s.client.TimeOut)

	if s.SessionToken != "" {
		req.Set(headerSessionToken, s.SessionToken)
	}
}

// Create a request which is set mater key
func (s *ParseSession) initMasterRequest(req *gorequest.SuperAgent) {
	req.
		Set(headerAppID, s.client.ApplicationID).
		Set(headerMasterKey, s.client.MasterKey).
		Timeout(s.client.TimeOut)

	if s.SessionToken != "" {
		req.Set(headerSessionToken, s.SessionToken)
	}
}

func (s *ParseSession) get(path string) *gorequest.SuperAgent {
	req := gorequest.New().Get(s.client.URL + path)
	s.initRequest(req)
	return req
}

func (s *ParseSession) getByMaster(path string) *gorequest.SuperAgent {
	req := gorequest.New().Get(s.client.URL + path)
	s.initMasterRequest(req)
	return req
}

func (s *ParseSession) post(path string) *gorequest.SuperAgent {
	req := gorequest.New().Post(s.client.URL + path)
	s.initRequest(req)
	return req
}

func (s *ParseSession) del(path string) *gorequest.SuperAgent {
	req := gorequest.New().Delete(s.client.URL + path)
	s.initRequest(req)
	return req
}

// Signup new user
func (s *ParseSession) Signup(data Signup) (user User, err error) {
	return user, do(s.post("/users").Send(data), &user)
}

// Login with data
func (s *ParseSession) Login(username string, password string) (user User, err error) {

	// Query values
	vals := url.Values{
		"username": []string{username},
		"password": []string{password},
	}

	// Create a user
	err = do(s.get("/login").Query(vals.Encode()), &user)

	if user.SessionToken != "" {
		s.SessionToken = user.SessionToken
	}

	return user, err
}

// GetUser gets user information
func (s *ParseSession) GetUser(userObjectID string) (user User, err error) {
	if userObjectID == "" {
		return user, errors.New("userObjectID must not be empty")
	}
	return user, do(s.get("/users/"+userObjectID), &user)
}

// GetUserByMaster gets user information
func (s *ParseSession) GetUserByMaster(userObjectID string) (user User, err error) {
	if userObjectID == "" {
		return user, errors.New("userObjectID must not be empty")
	}
	return user, do(s.getByMaster("/users/"+userObjectID), &user)
}

// GetMe gets self user information
func (s *ParseSession) GetMe() (user User, err error) {
	err = s.GetMeInto(&user)
	return user, err
}

// GetMeInto gets self user information into provided object
func (s *ParseSession) GetMeInto(user interface{}) error {
	if user == nil {
		return errors.New("user must not be nil")
	}
	return do(s.get("/users/me"), user)
}

// DeleteUser deletes user by ID
func (s *ParseSession) DeleteUser(userID string) error {
	return do(s.del("/users/"+userID), nil)
}

// UploadInstallation stores the subscription data for installations
func (s *ParseSession) UploadInstallation(data Installation, result interface{}) error {
	return do(s.post("/installations").Send(data), result)
}

// PushNotification sends push-notifiaction each device via parse
func (s *ParseSession) PushNotification(query map[string]interface{}, data PushNotificationData) error {
	body := PushNotificationQuery{
		Where: query,
		Data:  data,
	}
	return do(s.post("/push").Send(body), nil)
}

// Execute a parse request
func do(req *gorequest.SuperAgent, data interface{}) error {

	res, body, errs := req.End()
	if errs != nil {
		return fmt.Errorf("%v", errs)
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		// parse as error model
		reserr := new(Error)
		err := json.NewDecoder(strings.NewReader(body)).Decode(reserr)
		if err != nil {
			return err
		}
		return errors.New(reserr.Message + " - code:" + strconv.Itoa(reserr.Code))
	}
	if data == nil {
		return nil
	}
	return json.NewDecoder(strings.NewReader(body)).Decode(data)
}
