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

	})

}
