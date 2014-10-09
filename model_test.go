package goparse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"encoding/json"
	"strings"
	"time"
)

func TestModel(t *testing.T) {

	Convey("Start model testing", t, func() {

		s := `{
				"sessionToken":"2qx3lA4So1jDzW1Dgj2Fq3sCC",
				"objectId":"5J9g2dDJTF",
				"username":"test",
				"phone":"090-1234-5678",
				"createdAt":"2014-09-29T08:31:59.030Z",
				"updatedAt":"2014-10-29T08:31:59.030Z"
			  }`

		Convey("Decode JSON to User", func() {

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

	})

}
