package pantheon

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInSlice(t *testing.T) {
	Convey("Given I have a slice", t, func() {
		s := []string{"foo", "bar", "fizz"}

		Convey("When I call inSlice, passing in the slice and a value that is in the slice", func() {
			res := inSlice("foo", s)

			Convey("Then the result of inSlice will be true", func() {
				So(res, ShouldBeTrue)
			})
		})

		Convey("When I call inSlice, passing in the slice and a value that is NOT in the slice", func() {
			res := inSlice("not in slice", s)

			Convey("Then the result of inSlice will be false", func() {
				So(res, ShouldBeFalse)
			})
		})
	})
}
