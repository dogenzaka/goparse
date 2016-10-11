package goparse

import (
	"strconv"
	"time"
)

type (
	// User data type
	User struct {
		ObjectID      string    `json:"objectId,omitempty"`
		Email         string    `json:"email,omitempty"`
		UserName      string    `json:"username,omitempty"`
		Phone         string    `json:"phone,omitempty"`
		EmailVerified bool      `json:"emailVerified,omitempty"`
		SessionToken  string    `json:"sessionToken,omitempty"`
		AuthData      *AuthData `json:"authData,omitempty"`
		CreatedAt     time.Time `json:"createdAt,omitempty"`
		UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	}

	// AuthData data type
	AuthData struct {
		Facebook  *Facebook  `json:"facebook,omitempty"`
		Twitter   *Twitter   `json:"twitter,omitempty"`
		Anonymous *Anonymous `json:"anonymous,omitempty"`
	}

	// Facebook data type
	Facebook struct {
		ID          string    `json:"id,omitempty"`
		AccessToken string    `json:"access_token,omitempty"`
		Expiration  time.Time `json:"expiration_date,omitempty"`
	}

	// Twitter data type
	Twitter struct {
		ID              string `json:"id,omitempty"`
		ScreenName      string `json:"screen_name,omitempty"`
		ConsumerKey     string `json:"consumer_key,omitempty"`
		ConsumerSecret  string `json:"consumer_secret,omitempty"`
		AuthToken       string `json:"auth_token,omitempty"`
		AuthTokenSecret string `json:"auth_token_secret,omitempty"`
	}

	// Anonymous data type
	Anonymous struct {
		ID string `json:"id,omitempty"`
	}

	// Signup data type
	Signup struct {
		UserName string    `json:"username"`
		Password string    `json:"password"`
		AuthData *AuthData `json:"authData,omitempty"`
	}

	// Installation data type
	Installation struct {
		ObjectID       string    `json:"objectId,omitempty"`
		GCMSenderID    string    `json:"GCMSenderID,omitempty"`
		AppIdentifier  string    `json:"appIdentifier,omitempty"`
		AppName        string    `json:"appName,omitempty"`
		AppVersion     string    `json:"appVersion,omitempty"`
		Badge          string    `json:"badge,omitempty"`
		Channels       []string  `json:"channels,omitempty"`
		DeviceToken    string    `json:"deviceToken,omitempty"`
		DeviceType     string    `json:"deviceType,omitempty"`
		InstallationID string    `json:"installationId,omitempty"`
		ParseVersion   string    `json:"parseVersion,omitempty"`
		PushType       string    `json:"pushType,omitempty"`
		TimeZone       string    `json:"timeZone,omitempty"`
		User           Pointer   `json:"user,omitempty"`
		CreatedAt      time.Time `json:"createdAt,omitempty"`
		UpdatedAt      time.Time `json:"updatedAt,omitempty"`
	}

	// Pointer data type
	Pointer struct {
		Type      string `json:"__type"`
		ClassName string `json:"className"`
		ObjectID  string `json:"objectId"`
	}

	// PushNotificationQuery data type.
	// You can set the push_time and expiration_time to either "2015-08-022T12:00:00.000Z"
	// or 1440226800.
	PushNotificationQuery struct {
		Where              map[string]interface{} `json:"where"`
		PushTime           interface{}            `json:"push_time,omitempty"`
		ExpirationTime     interface{}            `json:"expiration_time,omitempty"`
		ExpirationInterval int64                  `json:"expiration_interval,omitempty"`
		Data               interface{}            `json:"data"`
	}

	// ObjectResponse data type
	ObjectResponse struct {
		ObjectID     string    `json:"objectId,omitempty"`
		SessionToken string    `json:"sessionToken,omitempty"`
		CreatedAt    time.Time `json:"createdAt,omitempty"`
		UpdatedAt    time.Time `json:"updatedAt,omitempty"`
	}

	// Error data type
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"error"`
	}
)

// Error to string
func (err *Error) Error() string {
	return err.Message + " - code:" + strconv.Itoa(err.Code)
}
