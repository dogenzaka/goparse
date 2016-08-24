package goparse

import (
	"encoding/json"
	"os"
	"strings"
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

		uuid := "90f37332-48ab-d5ec-1267-cbdb7bd4a480"

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
				AuthData: &AuthData{
					Anonymous: &Anonymous{
						ID: uuid,
					},
				},
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
				So(IsObjectNotFound(err), ShouldBeTrue)
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
						So(user2.AuthData, ShouldNotBeNil)
						So(user2.AuthData.Anonymous, ShouldNotBeNil)
						So(user2.AuthData.Anonymous.ID, ShouldEqual, uuid)
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

				Convey("When get user into provided object with valid values", func() {

					var user2 User
					err := session.GetUserInto(user.ObjectID, &user2)

					Convey("It returns no errors", func() {
						So(err, ShouldBeNil)
						So(user2.ObjectID, ShouldEqual, user.ObjectID)
						So(user2.UserName, ShouldEqual, user.UserName)
						So(user2.SessionToken, ShouldNotBeEmpty)
						So(user2.AuthData, ShouldNotBeNil)
						So(user2.AuthData.Anonymous, ShouldNotBeNil)
						So(user2.AuthData.Anonymous.ID, ShouldEqual, uuid)
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

					Convey("When uploading user data by use masterKey", func() {
						data := map[string]string{
							"phone": "03-1200-3400",
						}
						buf, _ := json.Marshal(data)
						resp, err := sessionInMaster.UpdateUserByMaster(user.ObjectID, string(buf))

						Convey("It returns no errors", func() {
							So(err, ShouldBeNil)
							So(resp, ShouldNotBeNil)
							So(resp.UpdatedAt.Unix(), ShouldBeGreaterThan, 0)
						})

						Convey("Check the user data", func() {
							me, err := session.GetUser(user.ObjectID)
							So(err, ShouldBeNil)
							So(me, ShouldNotBeNil)
							So(me.ObjectID, ShouldEqual, user.ObjectID)
							So(me.Phone, ShouldEqual, user.Phone)
						})

						Convey("masterKey is empty", func() {
							data := map[string]string{
								"phone": "03-1200-3400",
							}
							sessionInMaster.client.MasterKey = ""
							_, err := sessionInMaster.UpdateUserByMaster(user.ObjectID, data)
							So(err, ShouldNotBeNil)
						})
					})

					Convey("When get user with empty masterKey", func() {

						sessionInMaster.client.MasterKey = ""
						_, err := sessionInMaster.GetUserByMaster(user.ObjectID)

						Convey("It returns error", func() {
							So(err, ShouldNotBeNil)
						})
					})

					Convey("When get user into provided object with empty sessionToken", func() {

						var user2 User
						err := sessionInMaster.GetUserIntoByMaster(user.ObjectID, &user2)

						Convey("It returns no errors", func() {
							So(err, ShouldBeNil)
							So(user2.ObjectID, ShouldEqual, user.ObjectID)
							So(user2.UserName, ShouldEqual, user.UserName)
							So(user2.SessionToken, ShouldNotBeEmpty)
						})
					})

					Convey("When sending user push as master", func() {

						body := PushNotificationQuery{
							Where: map[string]interface{}{
								"DeviceType": "android",
							},
							Data: map[string]string{
								"alert": "master push",
							},
						}
						err := sessionInMaster.PushNotificationByMaster(body)

						Convey("It returns no errors", func() {
							So(err, ShouldBeNil)
						})

					})

				})

				Convey("Parse class operation", func() {

					type Testdata struct {
						Code int64  `json:"code,omitempty"`
						Name string `json:"name,omitempty"`
						*ObjectResponse
					}
					testingClass := session.NewClass("Testdata")

					Convey("It returns no errors", func() {
						So(testingClass, ShouldNotBeNil)
					})
				})

			})

			Convey("When uploading user data", func() {
				data := map[string]string{
					"phone": "03-1200-2300",
				}
				buf, _ := json.Marshal(data)
				resp, err := session.UpdateUser(user.ObjectID, string(buf))

				Convey("It returns no errors", func() {
					So(err, ShouldBeNil)
					So(resp, ShouldNotBeNil)
					So(resp.UpdatedAt.Unix(), ShouldBeGreaterThan, 0)
				})

				Convey("Check the user data", func() {
					me, err := session.GetMe()
					So(err, ShouldBeNil)
					So(me, ShouldNotBeNil)
					So(me.ObjectID, ShouldEqual, user.ObjectID)
					So(me.Phone, ShouldEqual, user.Phone)
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

		Convey("When sending push-notification", func() {
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

			body := PushNotificationQuery{
				Where: map[string]interface{}{
					"objectId": result.ObjectID,
				},
				Data: map[string]string{
					"alert": "test push",
				},
			}

			err = session.PushNotification(body)

			Convey("It returns no errors", func() {
				So(err, ShouldBeNil)
			})
		})
		Convey("When login for revocable session", func() {
			client, err := getDefaultClient()
			So(err, ShouldBeNil)
			client.RevocableSession = true
			session := client.NewSession("")
			user, err := session.Login("testuser", "testpass")
			So(err, ShouldBeNil)
			So(strings.HasPrefix(user.SessionToken, "r:"), ShouldBeTrue)
			Convey("When logout", func() {
				session := client.NewSession(user.SessionToken)
				err := session.Logout()

				Convey("It returns no errors", func() {
					So(err, ShouldBeNil)
				})

				Convey("Check the user data", func() {
					me, err := session.GetMe()
					So(err, ShouldNotBeNil)
					So(me, ShouldNotBeNil)
					So(me.ObjectID, ShouldBeEmpty)
				})
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
