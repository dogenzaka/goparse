package goparse

import (
	"fmt"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseClass(t *testing.T) {

	defaultClient = nil
	os.Setenv("PARSE_APPLICATION_ID", "")
	os.Setenv("PARSE_APPLICATION_ID", os.Getenv("TEST_PARSE_APPLICATION_ID"))
	os.Setenv("PARSE_REST_API_KEY", os.Getenv("TEST_PARSE_REST_API_KEY"))

	Convey("With a valid keys", t, func() {

		client, err := getDefaultClient()
		So(err, ShouldBeNil)
		So(client.ApplicationID, ShouldEqual, os.Getenv("TEST_PARSE_APPLICATION_ID"))
		So(client.RESTAPIKey, ShouldEqual, os.Getenv("TEST_PARSE_REST_API_KEY"))

		session, err := NewSession("")

		Convey("When logging in valid parameters", func() {

			session.Login("testuser", "testpass")

			Convey("Parse class operation", func() {

				type Testdata struct {
					Code int64  `json:"code,omitempty"`
					Name string `json:"name,omitempty"`
					Key  string `json:"key,omitempty"`
					*ObjectResponse
				}
				testingClass := session.NewClass("Testdata")

				Convey("It returns no errors", func() {
					So(testingClass, ShouldNotBeNil)
				})

				Convey("Create class object", func() {
					data := Testdata{
						Code: 201,
						Name: "apple",
						Key:  fmt.Sprint(time.Now().UnixNano()),
					}
					var result Testdata
					err := testingClass.Create(data, &result)

					Convey("Checking", func() {
						So(err, ShouldBeNil)
						So(result, ShouldNotBeNil)
						So(result.ObjectID, ShouldNotBeEmpty)
						So(result.CreatedAt, ShouldNotBeEmpty)

						Convey("Select object by ID", func() {
							var result2 Testdata
							err := testingClass.Select(result.ObjectID, &result2)

							Convey("Checking", func() {
								So(err, ShouldBeNil)
								So(result2, ShouldNotBeNil)
								So(result2.ObjectID, ShouldEqual, result.ObjectID)
								So(result2.Code, ShouldEqual, data.Code)
								So(result2.Name, ShouldEqual, data.Name)
								So(result2.CreatedAt, ShouldNotBeEmpty)
								So(result2.UpdatedAt, ShouldNotBeEmpty)
							})
						})

						Convey("Select object by query", func() {
							type resultList struct {
								Results []*Testdata `json:"results"`
							}
							var result2 resultList
							params := map[string]interface{}{
								"key": data.Key,
							}
							err := testingClass.SelectQuery(params, &result2)

							Convey("Checking", func() {
								So(err, ShouldBeNil)
								So(result2, ShouldNotBeNil)
								So(len(result2.Results), ShouldEqual, 1)
								So(result2.Results[0].ObjectID, ShouldEqual, result.ObjectID)
								So(result2.Results[0].Code, ShouldEqual, data.Code)
								So(result2.Results[0].Name, ShouldEqual, data.Name)
								So(result2.Results[0].CreatedAt, ShouldNotBeEmpty)
								So(result2.Results[0].UpdatedAt, ShouldNotBeEmpty)
							})
						})

						Convey("It is not found", func() {
							var result2 Testdata
							err := testingClass.Select("hoge", &result2)

							Convey("Checking", func() {
								So(err, ShouldNotBeNil)
								So(IsObjectNotFound(err), ShouldBeTrue)
							})
						})
					})
				})

				Convey("Update class object", func() {
					data := Testdata{
						Code: 202,
						Name: "melon",
					}
					var result Testdata
					err := testingClass.Create(data, &result)

					Convey("Create after updated", func() {
						So(err, ShouldBeNil)
						So(result, ShouldNotBeNil)

						data := Testdata{
							Code: 2020,
							Name: "super-melon",
						}
						var result2 Testdata
						err := testingClass.Update(result.ObjectID, data, &result2)

						Convey("Checking", func() {
							So(err, ShouldBeNil)
							So(result2, ShouldNotBeNil)
							So(result2.UpdatedAt, ShouldNotBeEmpty)
						})

						Convey("Select object", func() {
							var result3 Testdata
							err := testingClass.Select(result.ObjectID, &result3)

							Convey("Checking", func() {
								So(err, ShouldBeNil)
								So(result3, ShouldNotBeNil)
								So(result3.ObjectID, ShouldEqual, result.ObjectID)
								So(result3.Code, ShouldEqual, data.Code)
								So(result3.Name, ShouldEqual, data.Name)
							})
						})
					})
				})

				Convey("Delete class object", func() {
					data := Testdata{
						Code: 203,
						Name: "banana",
					}
					var result Testdata
					err := testingClass.Create(data, &result)

					Convey("Create after deleted", func() {
						So(err, ShouldBeNil)
						So(result, ShouldNotBeNil)

						err := testingClass.Delete(result.ObjectID)

						Convey("Checking", func() {
							So(err, ShouldBeNil)
						})

						Convey("Select object", func() {
							var result2 Testdata
							err := testingClass.Select(result.ObjectID, &result2)

							Convey("Checking", func() {
								So(err, ShouldNotBeNil)
							})
						})
					})
				})
			})

		})
	})

}
