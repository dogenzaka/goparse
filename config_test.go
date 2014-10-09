package goparse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {

	Convey("Testing config start", t, func() {

		Convey("Given all args", func() {
			c := ParseConfig{
				"test",
				"test",
				"test",
				1000,
			}
			err := checkConfig(c)
			So(err, ShouldBeNil)
		})

		Convey("If url is blank", func() {
			c := ParseConfig{
				"",
				"test",
				"test",
				1000,
			}
			err := checkConfig(c)
			So(err, ShouldNotBeNil)
		})

		Convey("If application id is blank", func() {
			c := ParseConfig{
				"test",
				"",
				"test",
				1000,
			}
			err := checkConfig(c)
			So(err, ShouldNotBeNil)
		})

		Convey("If rest api key is blank", func() {
			c := ParseConfig{
				"test",
				"test",
				"",
				1000,
			}
			err := checkConfig(c)
			So(err, ShouldNotBeNil)
		})

		Convey("If time out is zero", func() {
			c := ParseConfig{
				"test",
				"test",
				"test",
				0,
			}
			err := checkConfig(c)
			So(err, ShouldNotBeNil)
		})

		Convey("If time out is a negative number", func() {
			c := ParseConfig{
				"test",
				"test",
				"test",
				-1,
			}
			err := checkConfig(c)
			So(err, ShouldNotBeNil)
		})
	})

}
