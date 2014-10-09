package goparse

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParse(t *testing.T) {

	Convey("Start parse test", t, func() {

		Convey("If session token is blank", func() {
			_, err := GetMe("")
			So(err, ShouldNotBeNil)
		})

	})

}
