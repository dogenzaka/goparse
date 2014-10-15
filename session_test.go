package goparse

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

const (
	testApplicationId = "jhEPVJsjS8vylJQWFMDUhdF53XV2WfNYTLnPk5KA"
	testRestAPIKey    = "ZJea2ZeJF0E862gR6QIEmfqTyrdetltDwyYdLwOP"
)

func TestParseSession(t *testing.T) {

	defaultClient = nil
	os.Setenv("PARSE_APPLICATION_ID", "")

	Convey("When creating a session", t, func() {

		Convey("Without environment value", func() {

			session, err := NewSession("SESSION TOKEN")

			Convey("It should return an error", func() {
				So(session, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "client requires PARSE_APPLICATION_ID")
			})

		})

		Convey("With session token", func() {

			os.Setenv("PARSE_APPLICATION_ID", "APP_ID")

			session, err := NewSession("SESSION TOKEN")

			Convey("It should return a valid client", func() {
				So(err, ShouldBeNil)
				So(session, ShouldNotBeNil)
			})

		})

	})

	os.Setenv("PARSE_APPLICATION_ID", os.Getenv("TEST_PARSE_APPLICATION_ID"))
	os.Setenv("PARSE_REST_API_KEY", os.Getenv("TEST_PARSE_REST_API_KEY"))

	defaultClient = nil

	Convey("With a valid keys", t, func() {

		client, err := getDefaultClient()
		So(err, ShouldBeNil)
		So(client.ApplicationId, ShouldEqual, os.Getenv("TEST_PARSE_APPLICATION_ID"))
		So(client.RESTAPIKey, ShouldEqual, os.Getenv("TEST_PARSE_REST_API_KEY"))

		session, err := NewSession("")

		Convey("When signing up with empty values", func() {

			err := session.Signup(Signup{
				UserName: "",
				Password: "",
			})

			Convey("It returns an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "username missing - code:200")
			})

		})

		Convey("When signing up with valid values", func() {

			err := session.Signup(Signup{
				UserName: "testuser",
				Password: "testpass",
			})

			Convey("It returns no errors", func() {
				So(err, ShouldBeNil)
			})

		})

		Convey("When logging in invalid parameters", func() {

			user, err := session.Login("unknown", "password")

			Convey("It returns an error", func() {
				So(user.ObjectId, ShouldEqual, "")
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "invalid login parameters - code:101")
			})

		})

		Convey("When logging in valid parameters", func() {

			user, err := session.Login("testuser", "testpass")

			Convey("It returns an user with token", func() {
				So(err, ShouldBeNil)
				So(user.ObjectId, ShouldNotEqual, "")
				So(user.SessionToken, ShouldNotEqual, "")
				So(user.UserName, ShouldEqual, "testuser")
			})

		})

		Convey("When deleting a user", func() {

			user, err := session.Login("testuser", "testpass")
			So(err, ShouldBeNil)

			err = session.DeleteUser(user.ObjectId)

			Convey("It returns no errors", func() {
				So(err, ShouldBeNil)
			})

		})

	})

}
