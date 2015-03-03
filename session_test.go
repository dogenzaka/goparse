package goparse

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
		So(client.ApplicationID, ShouldEqual, os.Getenv("TEST_PARSE_APPLICATION_ID"))
		So(client.RESTAPIKey, ShouldEqual, os.Getenv("TEST_PARSE_REST_API_KEY"))

		session, err := NewSession("")

		Convey("When signing up with empty values", func() {

			_, err := session.Signup(Signup{
				UserName: "",
				Password: "",
			})

			Convey("It returns an error", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "missing username - code:200")
			})

		})

		Convey("When signing up with valid values", func() {

			user, err := session.Signup(Signup{
				UserName: "testuser",
				Password: "testpass",
			})

			Convey("It returns no errors", func() {
				So(err, ShouldBeNil)
				So(user.SessionToken, ShouldNotBeEmpty)
				So(user.ObjectID, ShouldNotBeEmpty)
			})

		})

		Convey("When logging in invalid parameters", func() {

			user, err := session.Login("unknown", "password")

			Convey("It returns an error", func() {
				So(user.ObjectID, ShouldEqual, "")
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "invalid login parameters - code:101")
			})

		})

		Convey("When logging in valid parameters", func() {

			user, err := session.Login("testuser", "testpass")

			Convey("It returns an user with token", func() {
				So(err, ShouldBeNil)
				So(user.ObjectID, ShouldNotEqual, "")
				So(user.SessionToken, ShouldNotEqual, "")
				So(user.UserName, ShouldEqual, "testuser")

				Convey("When get user with empty values", func() {

					_, err := session.GetUser("")

					Convey("It returns an error", func() {
						So(err, ShouldNotBeNil)
					})
				})

				Convey("When get user with valid values", func() {

					user2, err := session.GetUser(user.ObjectID)

					Convey("It returns no errors", func() {
						So(err, ShouldBeNil)
						So(user2.ObjectID, ShouldEqual, user.ObjectID)
						So(user2.UserName, ShouldEqual, user.UserName)
						So(user2.SessionToken, ShouldNotBeEmpty)
					})
				})

				Convey("When get user with empty sessionToken", func() {

					session.SessionToken = ""
					user2, err := session.GetUser(user.ObjectID)

					Convey("It returns no errors", func() {
						So(err, ShouldBeNil)
						So(user2.ObjectID, ShouldEqual, user.ObjectID)
						So(user2.UserName, ShouldEqual, user.UserName)
						So(user2.SessionToken, ShouldBeEmpty)
					})
				})

				Convey("Create client in master key", func() {

					os.Setenv("PARSE_MASTER_KEY", os.Getenv("TEST_PARSE_MASTER_KEY"))

					clientInMaster, err := NewClient()
					So(err, ShouldBeNil)
					So(clientInMaster.MasterKey, ShouldEqual, os.Getenv("TEST_PARSE_MASTER_KEY"))

					sessionInMaster := clientInMaster.NewSession("")

					Convey("When get user with empty values", func() {

						_, err := sessionInMaster.GetUserByMaster("")

						Convey("It returns an error", func() {
							So(err, ShouldNotBeNil)
						})
					})

					Convey("When get user with empty sessionToken", func() {

						user2, err := sessionInMaster.GetUserByMaster(user.ObjectID)

						Convey("It returns no errors", func() {
							So(err, ShouldBeNil)
							So(user2.ObjectID, ShouldEqual, user.ObjectID)
							So(user2.UserName, ShouldEqual, user.UserName)
							So(user2.SessionToken, ShouldNotBeEmpty)
						})
					})
				})

			})

		})

		Convey("When uploading installation data", func() {

			user, err := session.Login("testuser", "testpass")
			So(err, ShouldBeNil)

			data := Installation{
				AppName:        "Push",
				AppIdentifier:  "com.push.app",
				AppVersion:     "1.0",
				DeviceType:     "android",
				InstallationID: "8a779f48-0141-4dfa-ba5f-ac49c794efd5",
				ParseVersion:   "1.8.2",
				TimeZone:       "Asia/Tokyo",
				User: Pointer{
					Type:      "Pointer",
					ClassName: "_User",
					ObjectID:  user.ObjectID,
				},
			}
			result := Installation{}
			err = session.UploadInstallation(data, &result)

			Convey("It returns no errors", func() {
				So(err, ShouldBeNil)
			})

		})

		Convey("When sending push-notifiaction", func() {
			user, err := session.Login("testuser", "testpass")
			So(err, ShouldBeNil)

			installation := Installation{
				AppName:        "Push",
				AppIdentifier:  "com.push.app",
				AppVersion:     "1.0",
				DeviceType:     "android",
				InstallationID: "8a779f48-0141-4dfa-ba5f-ac49c794efd5",
				ParseVersion:   "1.8.2",
				TimeZone:       "Asia/Tokyo",
				User: Pointer{
					Type:      "Pointer",
					ClassName: "_User",
					ObjectID:  user.ObjectID,
				},
			}
			result := Installation{}
			err = session.UploadInstallation(installation, &result)
			So(err, ShouldBeNil)

			query := map[string]interface{}{
				"objectId": result.ObjectID,
			}
			data := PushNotificationData{
				Alert: "test push",
			}

			err = session.PushNotification(query, data)

			Convey("It returns no errors", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When deleting a user", func() {

			user, err := session.Login("testuser", "testpass")
			So(err, ShouldBeNil)

			err = session.DeleteUser(user.ObjectID)

			Convey("It returns no errors", func() {
				So(err, ShouldBeNil)
			})

		})

	})

}
