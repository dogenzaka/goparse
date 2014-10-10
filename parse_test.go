package goparse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {

	Convey("Start parse test", t, func() {

		Convey("If New with invalid config", func() {
			c := ParseConfig{}
			_, err := New(c)
			So(err, ShouldNotBeNil)
		})

		Convey("If session token is empty", func() {
			c := ParseConfig{"test", "test", "test", "test", 1000}
			pc, err := New(c)
			So(err, ShouldBeNil)
			_, err = pc.GetMe("")
			So(err, ShouldNotBeNil)
		})

		Convey("If session token is empty", func() {
			c := ParseConfig{"test", "test", "test", "test", 1000}
			pc, err := New(c)
			So(err, ShouldBeNil)
			err = pc.GetCustomMe("", new(User))
			So(err, ShouldNotBeNil)
		})

		Convey("If interface{} is empty", func() {
			c := ParseConfig{"test", "test", "test", "test", 1000}
			pc, err := New(c)
			So(err, ShouldBeNil)
			err = pc.GetCustomMe("test", nil)
			So(err, ShouldNotBeNil)
		})

	})

}
