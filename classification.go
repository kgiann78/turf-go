package turf

import (
	"github.com/tomchavakis/turf-go/geojson/geometry"
	"math"
)

// NearestPoint takes a reference point and a list of points and returns the point from the point list closest to the reference.
func NearestPoint(refPoint geometry.Point, points []geometry.Point) geometry.Point {
	if len(points) == 0 {
		return refPoint
	}

	result := points[0]
	minDist := math.MaxFloat64

	for _, point := range points {
		dist := PointDistance(refPoint, point)
		if dist < minDist {
			result = point
			minDist = dist
		}
	}

	return result
}
