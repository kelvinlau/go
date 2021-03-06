package geometry3d

import (
	. "github.com/kelvinlau/go/floats"
	"math"
)

// Ball is a sphere in 3D space.
type Ball struct {
	Point
	R float64
}

// IntersectionPointBallLine returns intersection points of a ball and a line,
// results may be of size 0, 1, or 2.
func IntersectionPointBallLine(b Ball, l Line) []Point {
	d := ClosestLinePoint(l, b.Point)
	e := Len2(Vec(b.Point, d))
	if e > Sqr(b.R) {
		return []Point{}
	}
	lv := LineVec(l)
	lv = Div(lv, Len(lv))
	t := math.Sqrt(Sqr(b.R) - e)
	if Sign(t) == 0 {
		return []Point{d}
	}
	v := Mul(lv, t)
	return []Point{Add(d, v), Sub(d, v)}
}

// GreatCircle is a function that given the latitude and longitude of two points
// in degrees, calculates the distance over the sphere between them.
// Latitude is given in the range [-pi/2, pi/2] degrees,
// Longitude is given in the range [-pi,pi] degrees.
func GreatCircle(lat1, long1, lat2, long2 float64) float64 {
	return math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(long2-long1))
}

// det solves the determinant of a matrix recursively.
func det(m [][]float64) (res float64) {
	n := len(m)
	if n == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0]
	}
	for skip := 0; skip < n; skip++ {
		s := [][]float64{}
		for i := 1; i < n; i++ {
			sr := []float64{}
			for j := 0; j < n; j++ {
				if j == skip {
					continue
				}
				sr = append(sr, m[i][j])
			}
			s = append(s, sr)
		}
		x := det(s)
		if skip%2 != 0 {
			res -= m[0][skip] * x
		} else {
			res += m[0][skip] * x
		}
	}
	return res
}

// Ball4 returns a ball passing 4 points.
func Ball4(ps [4]Point) *Ball {
	s := [][]float64{}
	for i := 0; i < 4; i++ {
		s = append(s, []float64{
			ps[i].X,
			ps[i].Y,
			ps[i].Z,
			1,
		})
	}
	if Sign(det(s)) == 0 {
		return nil
	}

	m := [][]float64{}
	for i := 0; i < 4; i++ {
		m = append(m, []float64{
			Sqr(ps[i].X) + Sqr(ps[i].Y) + Sqr(ps[i].Z),
			ps[i].X,
			ps[i].Y,
			ps[i].Z,
			1,
		})
	}
	sol := []float64{}
	for skip := 0; skip < 5; skip++ {
		for i := 0; i < 4; i++ {
			for j, sn := 0, 0; j < 5; j++ {
				if j == skip {
					continue
				}
				s[i][sn] = m[i][j]
				sn++
			}
		}
		sol = append(sol, det(s))
	}
	for i := 1; i < 5; i++ {
		if i%2 != 0 {
			sol[i] /= sol[0]
		} else {
			sol[i] /= -sol[0]
		}
	}
	for i := 1; i < 4; i++ {
		sol[i] /= 2
		sol[4] += Sqr(sol[i])
	}
	return &Ball{Point{sol[1], sol[2], sol[3]}, sol[4]}
}
