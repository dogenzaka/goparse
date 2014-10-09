package goparse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {

	Convey("Testing config start", t, func() {

		Convey("Given all args", func() {
			err := InitConfig("test", "test", "test", 1000)
			So(err, ShouldBeNil)
		})

		Convey("If url is blank", func() {
			err := InitConfig("", "test", "test", 1000)
			So(err, ShouldNotBeNil)
		})

		Convey("If application id is blank", func() {
			err := InitConfig("test", "", "test", 1000)
			So(err, ShouldNotBeNil)
		})

		Convey("If rest api key is blank", func() {
			err := InitConfig("test", "test", "", 1000)
			So(err, ShouldNotBeNil)
		})

		Convey("If time out is zero", func() {
			err := InitConfig("test", "test", "test", 0)
			So(err, ShouldNotBeNil)
		})

		Convey("If time out is a negative number", func() {
			err := InitConfig("test", "test", "test", -1)
			So(err, ShouldNotBeNil)
		})
	})

}
