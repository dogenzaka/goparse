package goparse

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	headerAppID            = "X-Parse-Application-Id"    // Parse Application ID
	headerMasterKey        = "X-Parse-Master-Key"        // Parse Master Key
	headerAPIKey           = "X-Parse-REST-API-Key"      // Parse REST API Key
	headerSessionToken     = "X-Parse-Session-Token"     // Parse Session Token
	headerRevocableSession = "X-Parse-Revocable-Session" // Parse Session Token

	pathMe = "/users/me"
)

// ParseSession is the client which has SessionToken as user authentication.
type ParseSession struct {
	client       *ParseClient
	SessionToken string
}

const (
	// Error code referrense at parse.com
	// http://www.parse.com/docs/dotnet/api/html/T_Parse_ParseException_ErrorCode.htm

	// errCodeObjectNotFound is object not found
	errCodeObjectNotFound = 101
)

var (
	// ErrObjectNotFound Error code indicating the specified object doesn't exist.
	ErrObjectNotFound = errors.New("object not found")
)

// Create a request which is set headers for Parse API
func (s *ParseSession) initRequest(req *gorequest.SuperAgent, useMaster bool) {
	if useMaster {
		req.
			Set(headerAppID, s.client.ApplicationID).
			Set(headerMasterKey, s.client.MasterKey).
			Timeout(s.client.TimeOut)
	} else {
		req.
			Set(headerAppID, s.client.ApplicationID).
			Set(headerAPIKey, s.client.RESTAPIKey).
			Timeout(s.client.TimeOut)
	}

	if s.client.RevocableSession {
		req.Set(headerRevocableSession, "1")
	}

	if s.SessionToken != "" {
		req.Set(headerSessionToken, s.SessionToken)
	}
}

func (s *ParseSession) get(path string, useMaster bool) *gorequest.SuperAgent {
	req := gorequest.New().Get(s.client.URL + path)
	s.initRequest(req, useMaster)
	return req
}

func (s *ParseSession) post(path string, useMaster bool) *gorequest.SuperAgent {
	req := gorequest.New().Post(s.client.URL + path)
	s.initRequest(req, useMaster)
	return req
}

func (s *ParseSession) put(path string, useMaster bool) *gorequest.SuperAgent {
	req := gorequest.New().Put(s.client.URL + path)
	s.initRequest(req, useMaster)
	return req
}

func (s *ParseSession) del(path string, useMaster bool) *gorequest.SuperAgent {
	req := gorequest.New().Delete(s.client.URL + path)
	s.initRequest(req, useMaster)
	return req
}

// Signup new user
func (s *ParseSession) Signup(data interface{}) (user User, err error) {
	return user, do(s.post("/users", false).Send(data), &user)
}

// Login with data
func (s *ParseSession) Login(username string, password string) (user User, err error) {

	// Query values
	vals := url.Values{
		"username": []string{username},
		"password": []string{password},
	}

	// Create a user
	err = do(s.get("/login", false).Query(vals.Encode()), &user)

	if user.SessionToken != "" {
		s.SessionToken = user.SessionToken
	}

	return user, err
}

// Logout deletes session from parse
func (s *ParseSession) Logout() (err error) {
	return do(s.post("/logout", false), nil)
}

// RequestPasswordReset let parse server to send a password reset mail
func (s *ParseSession) RequestPasswordReset(email string) (err error) {
	return do(s.post("/requestPasswordReset", false).Send(User{
		Email: email,
	}), nil)
}

// GetUser gets user information
func (s *ParseSession) GetUser(userObjectID string) (user User, err error) {
	return user, s.getUser(userObjectID, &user, false)
}

// GetUserByMaster gets user information by use master key
func (s *ParseSession) GetUserByMaster(userObjectID string) (user User, err error) {
	return user, s.getUser(userObjectID, &user, true)
}

// GetUserInto gets user information into provided object
func (s *ParseSession) GetUserInto(userObjectID string, user interface{}) (err error) {
	return s.getUser(userObjectID, user, false)
}

// GetUserIntoByMaster gets user information into provided object by use master key
func (s *ParseSession) GetUserIntoByMaster(userObjectID string, user interface{}) (err error) {
	return s.getUser(userObjectID, user, true)
}

// GetUserByMaster gets user information by private
func (s *ParseSession) getUser(userObjectID string, user interface{}, useMaster bool) (err error) {
	if userObjectID == "" {
		return errors.New("userObjectID must not be empty")
	}
	if useMaster && s.client.MasterKey == "" {
		return errors.New("request is requires PARSE_REST_API_KEY")
	}
	return do(s.get("/users/"+userObjectID, useMaster), &user)
}

// UpdateUser update user information
func (s *ParseSession) UpdateUser(userObjectID string, data interface{}) (*ObjectResponse, error) {
	return s.updateUser(userObjectID, data, false)
}

// UpdateUserByMaster update user information by use master key
func (s *ParseSession) UpdateUserByMaster(userObjectID string, data interface{}) (*ObjectResponse, error) {
	return s.updateUser(userObjectID, data, true)
}

// UpdateUser update user information by private
func (s *ParseSession) updateUser(userObjectID string, data interface{}, useMaster bool) (*ObjectResponse, error) {
	if userObjectID == "" {
		return nil, errors.New("userObjectID must not be empty")
	}
	if useMaster && s.client.MasterKey == "" {
		return nil, errors.New("request is requires PARSE_REST_API_KEY")
	}
	var resp ObjectResponse
	return &resp, do(s.put("/users/"+userObjectID, useMaster).Send(data), &resp)
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
	return do(s.get("/users/me", false), user)
}

// DeleteUser deletes user by ID
func (s *ParseSession) DeleteUser(userID string) error {
	return do(s.del("/users/"+userID, false), nil)
}

// UploadInstallation stores the subscription data for installations
func (s *ParseSession) UploadInstallation(data Installation, result interface{}) error {
	return do(s.post("/installations", false).Send(data), result)
}

// PushNotification sends push-notifiaction each device via parse
func (s *ParseSession) PushNotification(body PushNotificationQuery) error {
	return do(s.post("/push", false).Send(body), nil)
}

// PushNotificationByMaster sends push-notifiaction as Master to each device via parse
func (s *ParseSession) PushNotificationByMaster(body PushNotificationQuery) error {
	return do(s.post("/push", true).Send(body), nil)
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
		return reserr
	}
	if data == nil {
		return nil
	}
	return json.NewDecoder(strings.NewReader(body)).Decode(data)
}

// NewClass creates a new class from the session
func (s *ParseSession) NewClass(className string) *ParseClass {
	return &ParseClass{
		Session:   s,
		Name:      className,
		ClassURL:  "/classes/" + className,
		UseMaster: false,
	}
}

// IsObjectNotFound check the error "not found" or not
func IsObjectNotFound(err error) bool {
	v, ok := err.(*Error)
	return ok && v.Code == errCodeObjectNotFound
}
