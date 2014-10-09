package goparse

import (
	"testing"

	"encoding/json"
	"strings"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestModel(t *testing.T) {

	Convey("Start model testing", t, func() {

		Convey("Decode JSON to User", func() {

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

			Convey("Exsits SessionToken", func() {
				So(u.SessionToken, ShouldEqual, "2qx3lA4So1jDzW1Dgj2Fq3sCC")
			})

			Convey("Exsits ObjectId", func() {
				So(u.ObjectId, ShouldEqual, "5J9g2dDJTF")
			})

			Convey("Exsits UserName", func() {
				So(u.UserName, ShouldEqual, "test")
			})

			Convey("Exsits Phone", func() {
				So(u.Phone, ShouldEqual, "090-1234-5678")
			})

			Convey("Exsits CreatedAt", func() {
				t, _ := time.Parse(time.RFC3339, "2014-09-29T08:31:59.030Z")
				So(u.CreatedAt, ShouldHappenOnOrBefore, t)
			})

			Convey("Exsits UpdatedAt", func() {
				t, _ := time.Parse(time.RFC3339, "2014-10-29T08:31:59.030Z")
				So(u.UpdatedAt, ShouldHappenOnOrBefore, t)
			})
		})

		Convey("Decode JSON to Facebook", func() {

			s := `{
				"id": "abcdefg",
				"access_token": "hijklmn",
				"expiration_date": "2014-10-29T08:31:59.030Z"
			 }`

			var f Facebook
			err := json.NewDecoder(strings.NewReader(s)).Decode(&f)
			So(err, ShouldEqual, nil)

			Convey("Exsits facebook id", func() {
				So(f.Id, ShouldEqual, "abcdefg")
			})

			Convey("Exsits access_token", func() {
				So(f.AccessToken, ShouldEqual, "hijklmn")
			})

			Convey("Exsits expiration_date", func() {
				t, _ := time.Parse(time.RFC3339, "2014-10-29T08:31:59.030Z")
				So(f.Expiration, ShouldHappenOnOrBefore, t)
			})
		})

		Convey("Decode JSON to twitter", func() {

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

			Convey("Exsits twitter Id", func() {
				So(t.Id, ShouldEqual, "abcdefg")
			})

			Convey("Exsits screen_name", func() {
				So(t.ScreenName, ShouldEqual, "hijklmn")
			})

			Convey("Exsits consumer_key", func() {
				So(t.ConsumerKey, ShouldEqual, "123abc")
			})

			Convey("Exsits consumer_secret", func() {
				So(t.ConsumerSecret, ShouldEqual, "345def")
			})

			Convey("Exsits auth_token", func() {
				So(t.AuthToken, ShouldEqual, "678ghi")
			})

			Convey("Exsits UpdatedAt", func() {
				So(t.AuthTokenSecret, ShouldEqual, "901jkl")
			})
		})

		Convey("Decode JSON to Facebook", func() {

			s := `{
					"id": "abcdefg"
			  }`

			var a Anonymous
			err := json.NewDecoder(strings.NewReader(s)).Decode(&a)
			So(err, ShouldEqual, nil)

			Convey("Exsits facebook id", func() {
				So(a.Id, ShouldEqual, "abcdefg")
			})
		})

	})
}
