package goparse

import "time"

type (
	User struct {
		ObjectId      string    `json:"objectId,omitempty"`
		Email         string    `json:"email,omitempty"`
		Username      string    `json:"username,omitempty"`
		Phone         string    `json:"phone,omitempty"`
		EmailVerified bool      `json:"emailVerified,omitempty"`
		SessionToken  string    `json:"sessionToken,omitempty"`
		CreatedAt     time.Time `json:"createdAt,omitempty"`
		UpdatedAt     time.Time `json:"updatedAt,omitempty"`
	}

	AuthData struct {
		Facebook  *Facebook  `json:"facebook",omitempty`
		Twitter   *Twitter   `json:"twitter,omitempty"`
		Anonymous *Anonymous `json:"anonymous,omitempty"`
	}

	Facebook struct {
		ID          string    `json:"id,omitempty"`
		AccessToken string    `json:"access_token,omitempty"`
		Expiration  time.Time `json:"expiration_date,omitempty"`
	}

	Twitter struct {
		ID              string `json:"id,omitempty"`
		ScreenName      string `json:"screen_name,omitempty"`
		ConsumerKey     string `json:"consumer_key,omitempty"`
		ConsumerSecret  string `json:"consumer_secret,omitempty"`
		AuthToken       string `json:"auth_token,omitempty"`
		AuthTokenSecret string `json:"auth_token_secret,omitempty"`
	}

	Anonymous struct {
		ID string `json:"id,omitempty"`
	}

	ErrorReuslt struct {
		Code  int
		Error string
	}
)
