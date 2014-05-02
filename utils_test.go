// History: May 02 14 tcolar Creation

package utils

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {

	Convey("UnderscoredToSnake", t, func() {
		So(UnderscoredToSnake("this_is_an_example"), ShouldEqual, "ThisIsAnExample")
	})

	Convey("SnakeToUnderscore", t, func() {
		So(SnakeToUnderscore("ThisIsAnExample"), ShouldEqual, "this_is_an_example")
	})

	Convey("IsEmptyString", t, func() {
		So(IsEmptyString(""), ShouldEqual, true)
		So(IsEmptyString("a"), ShouldEqual, false)
	})
}

func TestNumber(t *testing.T) {

	Convey("StrToFloat", t, func() {
		f, _ := StrToFloat64("1.23")
		So(f, ShouldEqual, 1.23)
		f, _ = StrToFloat64("")
		So(f, ShouldEqual, 0.0)
	})

	Convey("RoundFloat", t, func() {
		f := 1.234567
		f2 := -95.23
		f3 := 9.98
		f4 := -9.98123456789
		So(RoundFloat(f, 0), ShouldEqual, 1)
		So(RoundFloat(f2, 0), ShouldEqual, -95)
		So(RoundFloat(f3, 0), ShouldEqual, 10)
		So(RoundFloat(f4, 0), ShouldEqual, -10)
		So(RoundFloat(f, 2), ShouldEqual, 1.23)
		So(RoundFloat(f2, 2), ShouldEqual, -95.23)
		So(RoundFloat(f3, 2), ShouldEqual, 9.98)
		So(RoundFloat(f4, 2), ShouldEqual, -9.98)
		So(RoundFloat(f, 8), ShouldEqual, 1.234567)
		So(RoundFloat(f2, 8), ShouldEqual, -95.23)
		So(RoundFloat(f3, 8), ShouldEqual, 9.98)
		So(RoundFloat(f4, 8), ShouldEqual, -9.98123457)
	})
}

func TestCollection(t *testing.T) {
	c1 := []string{}
	c2 := []string{"aa", "a", "d", "v"}
	c3 := []int{}
	c4 := []int{9, 7, 3, 5}

	Convey("HasString", t, func() {
		So(HasString(c1, "X"), ShouldEqual, false)
		So(HasString(c2, "v"), ShouldEqual, true)
	})

	Convey("FindString", t, func() {
		So(FindString(c1, "X"), ShouldEqual, -1)
		So(FindString(c2, "a"), ShouldEqual, 1)
	})

	Convey("HasInt", t, func() {
		So(HasInt(c3, -1), ShouldEqual, false)
		So(HasInt(c4, 9), ShouldEqual, true)
	})

	Convey("FindInt", t, func() {
		So(FindInt(c3, 99), ShouldEqual, -1)
		So(FindInt(c4, 9), ShouldEqual, 0)
	})
}
