package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTowSum(t *testing.T) {

	a := []int{2, 3, 6, 11, 15}
	b := []int{1, 2}
	Convey("将两数相加", t, func() {
		So(twoSum(a, 9), ShouldEqual, b)
	})

}
