package goparse

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseClient(t *testing.T) {

	Convey("When creating a client", t, func() {

		Convey("With no environment variables", func() {

			os.Setenv("PARSE_APPLICATION_ID", "")

			client, err := NewClient()

			Convey("It should return an error", func() {
				So(client, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "client requires PARSE_APPLICATION_ID")
			})

		})

		Convey("After setting PARSE_APPLICATION_ID", func() {

			os.Setenv("PARSE_APPLICATION_ID", "TEST_APP_ID")
			os.Setenv("PARSE_REST_API_KEY", "")

			client, err := NewClient()

			Convey("It should return an error", func() {
				So(client, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "client requires PARSE_REST_API_KEY")
			})

		})

		Convey("After setting PARSE_REST_API_KEY", func() {

			os.Setenv("PARSE_APPLICATION_ID", "TEST_APP_ID")
			os.Setenv("PARSE_REST_API_KEY", "TEST_API_KEY")

			client, err := NewClient()

			Convey("It should return a valid client", func() {
				So(client, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(client.ApplicationID, ShouldEqual, "TEST_APP_ID")
				So(client.URL, ShouldEqual, defaultEndPoint)
				So(client.RESTAPIKey, ShouldEqual, "TEST_API_KEY")
				So(client.MasterKey, ShouldEqual, "")
				So(client.TimeOut.Seconds(), ShouldEqual, 5)
			})

		})

		Convey("Without REST API KEY (custom end point)", func() {

			os.Setenv("PARSE_APPLICATION_ID", "TEST_APP_ID")
			os.Setenv("PARSE_ENDPOINT_URL", "https://dummy.com/")
			os.Setenv("PARSE_REST_API_KEY", "")

			client, err := NewClient()

			Convey("It should return a valid client", func() {
				os.Setenv("PARSE_ENDPOINT_URL", "")
				So(client, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(client.ApplicationID, ShouldEqual, "TEST_APP_ID")
				So(client.URL, ShouldEqual, "https://dummy.com/")
				So(client.RESTAPIKey, ShouldEqual, "")
				So(client.MasterKey, ShouldEqual, "")
				So(client.TimeOut.Seconds(), ShouldEqual, 5)
			})

		})
		Reset(func() {

			defaultClient = nil
			os.Setenv("PARSE_APPLICATION_ID", "")
			os.Setenv("PARSE_APPLICATION_ID", os.Getenv("TEST_PARSE_APPLICATION_ID"))
			os.Setenv("PARSE_REST_API_KEY", os.Getenv("TEST_PARSE_REST_API_KEY"))

		})
	})

	Convey("When creating a client with configuration", t, func() {

		Convey("Without empty configuration", func() {

			client, err := NewClientWithConfig(ParseConfig{})

			Convey("It should return an error", func() {
				So(client, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "client requires ApplicationId")
			})

		})

		Convey("Without REST API KEY if using (default end point)", func() {

			client, err := NewClientWithConfig(ParseConfig{
				ApplicationID: "APPID",
			})

			Convey("It should return an error", func() {
				So(client, ShouldBeNil)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "client requires RESTAPIKey")
			})

		})

		Convey("Without REST API KEY (custom end point)", func() {

			client, err := NewClientWithConfig(ParseConfig{
				ApplicationID: "APPID",
				URL:           "https://dummy.com/",
			})

			Convey("It should return a valid client", func() {
				So(client, ShouldNotBeNil)
				So(err, ShouldBeNil)
				So(client.ApplicationID, ShouldEqual, "APPID")
				So(client.URL, ShouldEqual, "https://dummy.com/")
				So(client.RESTAPIKey, ShouldEqual, "")
				So(client.MasterKey, ShouldEqual, "")
				So(client.TimeOut.Seconds(), ShouldEqual, 5)
			})

		})

		Convey("With valid config", func() {

			client, err := NewClientWithConfig(ParseConfig{
				ApplicationID: "APPID",
				RESTAPIKey:    "APIKEY",
			})

			Convey("It should return a valid client", func() {
				So(client, ShouldNotBeNil)
				So(err, ShouldBeNil)
			})

		})

	})

}
