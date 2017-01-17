package rocket

import (
	"math"
	"math/rand"

	"github.com/JamesRCooper/NeuralRockets/model"
	neural "github.com/JamesRCooper/geneticNeuralNetwork/model"
)

var numOfRockets = 80

var goalPosition = model.Vec{X: 350, Y: 50}
var startPosition = model.Vec{X: 350, Y: 300}

//MaxFlightTime designates how long a simulation will run before restarting
var MaxFlightTime = 750
var char = neural.CellCharacter{
	MutationRate: 0.0125, Activater: sigmoid, GeneCreator: geneCreator}

//InitRockets creates an array of pointers towards a new set of rockets for
//testing
func InitRockets() []*Rocket {
	builder := CreateBuilder(
		10,
		[]int{8, 6, 4, 2},
		&char,
		&goalPosition,
		startPosition,
		checkBoundry)

	rockets := make([]*Rocket, numOfRockets)
	for index := 0; index < numOfRockets; index++ {
		rockets[index] = builder.Build()
	}

	return rockets
}

func sigmoid(operand float64) float64 {
	return (2.0 / (1.0 + math.Exp(-1.0*operand))) - 1.0
}

func geneCreator() float64 {
	return 2.0*rand.Float64() - 1.0
}

func checkBoundry(p model.Vec) bool {
	if p.X < 0 || p.X > 700 {
		return false
	}

	if p.Y < 0 || p.Y > 400 {
		return false
	}

	if p.Y < 220 && p.Y > 180 && p.X < 500 && p.X > 200 {
		return false
	}

	return true
}
