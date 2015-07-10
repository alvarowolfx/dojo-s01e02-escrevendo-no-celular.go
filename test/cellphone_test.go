package test

import (
	. "dojo-s01e02-escrevendo-no-celular.go/lib"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCellphone(t *testing.T) {
	var cellphone *Cellphone

	Convey("Cellphone", t, func() {

		cellphone = NewCellphone()

		Convey("should return number representation when a letter is passed", func() {

			So(cellphone.ConvertToKeypadCode("A"), ShouldEqual, "2")
			So(cellphone.ConvertToKeypadCode("B"), ShouldEqual, "22")
			So(cellphone.ConvertToKeypadCode("C"), ShouldEqual, "222")

			So(cellphone.ConvertToKeypadCode("D"), ShouldEqual, "3")
			So(cellphone.ConvertToKeypadCode("E"), ShouldEqual, "33")
			So(cellphone.ConvertToKeypadCode("F"), ShouldEqual, "333")

			So(cellphone.ConvertToKeypadCode("P"), ShouldEqual, "7")
			So(cellphone.ConvertToKeypadCode("Q"), ShouldEqual, "77")
			So(cellphone.ConvertToKeypadCode("R"), ShouldEqual, "777")
			So(cellphone.ConvertToKeypadCode("S"), ShouldEqual, "7777")
		})

		Convey(" should use a underline to separate two number from the same key group", func() {

			So(cellphone.ConvertToKeypadCode("AB"), ShouldEqual, "2_22")
			So(cellphone.ConvertToKeypadCode("EF"), ShouldEqual, "33_333")
			So(cellphone.ConvertToKeypadCode("MAA"), ShouldEqual, "62_2")
		})

		Convey("should return '77773367_7773302_222337777_777766606660366656667889999_9999555337777' when 'SEMPRE ACESSO O DOJOPUZZLESSEMPRE ACESSO O DOJOPUZZLES' passed", func() {
			So(cellphone.ConvertToKeypadCode("SEMPRE ACESSO O DOJOPUZZLES"), ShouldEqual, "77773367_7773302_222337777_777766606660366656667889999_9999555337777")
		})

		Convey("should return '29337777666_63307777222_25552' when 'AWESOME GO' passed", func() {
			So(cellphone.ConvertToKeypadCode("AWESOME GO"), ShouldEqual, "29337777666_63304666")
		})

	})
}
