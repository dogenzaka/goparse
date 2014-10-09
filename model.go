package goparse

import "time"

type User struct {
	SessionToken string    `json:"sessionToken"`
	ObjectId     string    `json:"objectId"`
	UserName     string    `json:"username"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
