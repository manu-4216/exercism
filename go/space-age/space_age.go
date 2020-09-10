// Package space handles space time
package space

import "math"

// Planet type
type Planet string

const secodsInEarthYear = 31557600

var mapPlanets = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// round2Decimal rounds a number to 2 decimals
func round2Decimal(number float64) float64 {
	return math.Round(number*100) / 100
}

// Age converts the age on Earth in seconds to the age in years on a different planet
func Age(seconds float64, planet Planet) float64 {
	return round2Decimal(seconds / secodsInEarthYear / mapPlanets[planet])
}
