package valueobject

import (
	"fmt"
	"math"
)

const (
	maxLat        = 90.0
	maxLong       = 180.0
	earthRadiusKm = 6371.0
)

type (
	Location struct {
		lat  float64
		long float64
	}

	Locations []Location
)

func NewLocation(lat, long float64) (Location, error) {
	if lat < -maxLat || lat > maxLat {
		return Location{}, fmt.Errorf("latitude must be between %.2f and %.2f degrees", -maxLat, maxLat)
	}

	if long < -maxLong || long > maxLong {
		return Location{}, fmt.Errorf("longitude must be between %.2f and %.2f degrees", -maxLong, maxLong)
	}

	return Location{
		lat:  lat,
		long: long,
	}, nil
}

func (l Location) Lat() float64 {
	return l.lat
}

func (l Location) Long() float64 {
	return l.long
}

func (rcv Location) DistanceInKM(target Location) float64 {
	lat1 := rcv.lat * math.Pi / 180.0
	long1 := rcv.long * math.Pi / 180.0
	lat2 := target.lat * math.Pi / 180.0
	long2 := target.long * math.Pi / 180.0

	deltaLat := lat2 - lat1
	deltaLong := long2 - long1

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(deltaLong/2)*math.Sin(deltaLong/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}

func (rcv Location) IsInRadius(target Location, radius float64) bool {
	distance := rcv.DistanceInKM(target)

	return distance <= radius
}

func (rcv Locations) IsInRadius(target Location, radius float64) bool {
	for _, l := range rcv {
		if !l.IsInRadius(target, radius) {
			return false
		}
	}

	return true
}
