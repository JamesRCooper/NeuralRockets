package rocket

import (
	"math"
	"math/rand"
	"sort"

	"github.com/jamescooper/neuralRockets/model"
)

type rocketError struct {
	rocket *Rocket
	err    float64
}

//ByErr implements the required interface for running a sorting algorithm
type ByErr []rocketError

func (a ByErr) Len() int           { return len(a) }
func (a ByErr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByErr) Less(i, j int) bool { return a[i].err < a[j].err }

//Breed creates a breeding pool using the supplied slice of rockets, and breeds
//each rocket with a weighted random rocket from the pool. The weight of a rocket
//(number of times it appears in the pool) coincides to the total error of the
//rocket in the last run. The rocket with the highest error appears once in the
//pool; the rocket with the lowest error appears as many times as there are
//rockets, with a linear regression for all rockets in between.
func Breed(rockets []*Rocket) {
	rocketErrors := buildSortedErrors(rockets)
	for index := range rockets {
		rndValue := ((1.0 + rand.Float64()) / 2.0) * float64(numOfRockets)
		selection := int((1.0 + math.Sqrt(1.0+(8.0*rndValue))) / 2.0)
		partnerRocket := rocketErrors[selection]
		rockets[index].breed(*partnerRocket.rocket)

		rockets[index].position = startPosition
		rockets[index].velocity = model.Vec{X: 0, Y: 0}
		rockets[index].acceleration = model.Vec{X: 0, Y: 0}
		rockets[index].HitBoundry = false
		rockets[index].FlightTime = 0
	}
}

func buildSortedErrors(rockets []*Rocket) []rocketError {

	rocketErrors := make([]rocketError, numOfRockets)
	for index := 0; index < numOfRockets; index++ {
		currentError := processError(rockets[index])
		rocketErrors[index] = rocketError{rockets[index], currentError}
	}

	sort.Sort(ByErr(rocketErrors))
	return rocketErrors
}

func processError(r *Rocket) float64 {

	totalErr := 0.0

	rMapping := r.BuildMapping()
	xErr := math.Pow(rMapping.X-goalPosition.X, 2.0)
	yErr := math.Pow(rMapping.Y-goalPosition.Y, 2.0)
	totalErr += math.Sqrt(xErr + yErr)

	if r.HitBoundry {
		totalErr += 200.0 * (8.0 - (float64(r.FlightTime) / float64(MaxFlightTime)))
	}

	return totalErr
}
