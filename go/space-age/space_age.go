package space

// Planet represents a planet
type Planet string

const earthOrbitalPeriod float64 = 31557600

var planetOrbitalPeriod = map[Planet]float64{
	"Earth":   earthOrbitalPeriod,
	"Mercury": 0.2408467 * earthOrbitalPeriod,
	"Venus":   0.61519726 * earthOrbitalPeriod,
	"Mars":    1.8808158 * earthOrbitalPeriod,
	"Jupiter": 11.862615 * earthOrbitalPeriod,
	"Saturn":  29.447498 * earthOrbitalPeriod,
	"Uranus":  84.016846 * earthOrbitalPeriod,
	"Neptune": 164.79132 * earthOrbitalPeriod,
}

// Age calculates a person's age in a particular planet
func Age(ageInSeconds float64, planet Planet) float64 {
	return ageInSeconds / planetOrbitalPeriod[planet]
}
