package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReverse(t *testing.T) {
	Convey("反转整数", t, func() {
		So(reverse(123), ShouldEqual, 321)
	})
}
