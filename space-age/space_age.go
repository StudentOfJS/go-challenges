// Package space provides a function Age
package space

const earthYear float64 = 31557600

type Planet string

var yearLengths = map[Planet]float64{
	"Earth":   earthYear,
	"Mercury": 0.2408467 * earthYear,
	"Venus":   0.61519726 * earthYear,
	"Mars":    1.8808158 * earthYear,
	"Jupiter": 11.862615 * earthYear,
	"Saturn":  29.447498 * earthYear,
	"Uranus":  84.016846 * earthYear,
	"Neptune": 164.79132 * earthYear,
}

// Age takes age in seconds and a planet. Then returns your age on that planet
func Age(secs float64, planet Planet) float64 {
	return secs / yearLengths[planet]
}
