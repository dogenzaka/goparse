package goparse

import (
	"testing"

	"encoding/json"
	"strings"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestModel(t *testing.T) {

	Convey("Given JSON strings", t, func() {

		Convey("When decoding as User", func() {

			s := `{
				"sessionToken":"2qx3lA4So1jDzW1Dgj2Fq3sCC",
				"objectId":"5J9g2dDJTF",
				"username":"test",
				"phone":"090-1234-5678",
				"createdAt":"2014-09-29T08:31:59.030Z",
				"updatedAt":"2014-10-29T08:31:59.030Z"
			  }`

			var u User
			err := json.NewDecoder(strings.NewReader(s)).Decode(&u)
			So(err, ShouldEqual, nil)

			Convey("It has a SessionToken", func() {
				So(u.SessionToken, ShouldEqual, "2qx3lA4So1jDzW1Dgj2Fq3sCC")
			})

			Convey("It has an ObjectId", func() {
				So(u.ObjectId, ShouldEqual, "5J9g2dDJTF")
			})

			Convey("It has an UserName", func() {
				So(u.UserName, ShouldEqual, "test")
			})

			Convey("It has a Phone", func() {
				So(u.Phone, ShouldEqual, "090-1234-5678")
			})

			Convey("CreatedAt equals provided time", func() {
				t, _ := time.Parse(time.RFC3339, "2014-09-29T08:31:59.030Z")
				So(u.CreatedAt, ShouldHappenOnOrBefore, t)
			})

			Convey("UpdatedAt equals provided time", func() {
				t, _ := time.Parse(time.RFC3339, "2014-10-29T08:31:59.030Z")
				So(u.UpdatedAt, ShouldHappenOnOrBefore, t)
			})
		})

		Convey("When decoding as Facebook", func() {

			s := `{
				"id": "abcdefg",
				"access_token": "hijklmn",
				"expiration_date": "2014-10-29T08:31:59.030Z"
			 }`

			var f Facebook
			err := json.NewDecoder(strings.NewReader(s)).Decode(&f)
			So(err, ShouldEqual, nil)

			Convey("It has an Id", func() {
				So(f.Id, ShouldEqual, "abcdefg")
			})

			Convey("It has an AccessToken", func() {
				So(f.AccessToken, ShouldEqual, "hijklmn")
			})

			Convey("Expiration is provided date", func() {
				t, _ := time.Parse(time.RFC3339, "2014-10-29T08:31:59.030Z")
				So(f.Expiration, ShouldHappenOnOrBefore, t)
			})
		})

		Convey("When decoding as Twitter", func() {

			s := `{
				"id": "abcdefg",
				"screen_name": "hijklmn",
				"consumer_key": "123abc",
				"consumer_secret": "345def",
				"auth_token": "678ghi",
				"auth_token_secret": "901jkl"
			}`

			var t Twitter
			err := json.NewDecoder(strings.NewReader(s)).Decode(&t)
			So(err, ShouldEqual, nil)

			Convey("It has an Id", func() {
				So(t.Id, ShouldEqual, "abcdefg")
			})

			Convey("It has a ScreenName", func() {
				So(t.ScreenName, ShouldEqual, "hijklmn")
			})

			Convey("It has a ConsumerKey", func() {
				So(t.ConsumerKey, ShouldEqual, "123abc")
			})

			Convey("It has a ConsumerSecret", func() {
				So(t.ConsumerSecret, ShouldEqual, "345def")
			})

			Convey("It has an AuthToken", func() {
				So(t.AuthToken, ShouldEqual, "678ghi")
			})

			Convey("It has an AuthTokenSecret", func() {
				So(t.AuthTokenSecret, ShouldEqual, "901jkl")
			})
		})

		Convey("When decoding as Anonymous", func() {

			s := `{
					"id": "abcdefg"
			  }`

			var a Anonymous
			err := json.NewDecoder(strings.NewReader(s)).Decode(&a)
			So(err, ShouldEqual, nil)

			Convey("It has an Id", func() {
				So(a.Id, ShouldEqual, "abcdefg")
			})
		})

		Convey("When deocding as Error", func() {

			s := `{
				"code": 105,
				"error": "invalid field name: bl!ng"
			}`

			var a Error
			err := json.NewDecoder(strings.NewReader(s)).Decode(&a)
			So(err, ShouldEqual, nil)

			Convey("It has a Code", func() {
				So(a.Code, ShouldEqual, 105)
			})

			Convey("It has a Message", func() {
				So(a.Message, ShouldEqual, "invalid field name: bl!ng")
			})

		})

	})
}
