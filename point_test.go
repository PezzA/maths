package maths

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {
	Convey("Given the points of Pacman, inky and clide", t, func() {
		p := Point2D{0, -1}
		i := Point2D{1, 1}
		c := Point2D{2, -1}

		Convey("Should be able to apply a vector to pacman", func() {
			p := Point2D{1, 0}
			v := Vector2D{2, 3}
			So(p.AddVector(v), ShouldResemble, Point2D{3, 3})
		})

		Convey("Should be able to get the vector pacman and inky", func() {
			So(i.GetVectorToPoint(p), ShouldResemble, Vector2D{-1, -2})
		})

		Convey("should be able to get the distance between pacman and inky", func() {
			v := i.GetVectorToPoint(p)
			So(v.GetLength(), ShouldAlmostEqual, 2.23606797749979)
		})

		Convey("Should be able to work out if Inky or Clide is closer to pacman", func() {
			ip := i.GetVectorToPoint(p).GetLengthSquared()
			cp := c.GetVectorToPoint(p).GetLengthSquared()

			So(ip, ShouldEqual, 5)
			So(cp, ShouldEqual, 4)
		})

		Convey("Should be able to adjust pacman's speed", func() {
			v := Vector2D{3, 4}

			doubled := v.MultiplyBy(2)
			halved := v.DivideBy(2)

			So(v.GetLength(), ShouldEqual, 5)
			So(doubled.GetLength(), ShouldEqual, 10)
			So(halved.GetLength(), ShouldEqual, 2.5)
		})

		Convey("Should be able to normalise a vector", func() {
			i := Point2D{3, 4}
			p := Point2D{1, 2}

			v := p.GetVectorToPoint(i)

			So(v.Normalised(), ShouldResemble, Vector2D{0.7071067811865475, 0.7071067811865475})
			So(v.Normalised().GetLength(), ShouldAlmostEqual, 1)
		})

		Convey("Should be able to add two vectors together", func() {
			r := Vector2D{4, 0}
			d := Vector2D{0, -5}

			v := r.Add(d)

			So(v, ShouldResemble, Vector2D{4, -5})
		})

		Convey("Should be able to tell if 2 vectors are pointing in the same direction.", func() {
			r := Vector2D{2, 0}.Normalised()
			d := Vector2D{3, 0}.Normalised()

			So(r.dotProduct(d), ShouldEqual, 1)
		})

		Convey("Should be able to tell if 2 vectors are pointing in the opposite direction.", func() {
			r := Vector2D{-2, 0}.Normalised()
			d := Vector2D{3, 0}.Normalised()

			So(r.dotProduct(d), ShouldEqual, -1)
		})

		Convey("Should be able to tell if 2 vectors are pointing at right angles", func() {
			r := Vector2D{2, 0}.Normalised()
			d := Vector2D{0, 2}.Normalised()

			So(r.dotProduct(d), ShouldEqual, 0)
		})
	})
}
