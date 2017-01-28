package rocket

import (
	"github.com/JamesRCooper/NeuralRockets/model"
	m "github.com/JamesRCooper/geneticNeuralNetwork/model"
	"github.com/JamesRCooper/geneticNeuralNetwork/mutation"
)

var numOfRockets = 40

var goalPosition = model.Vec{X: 350, Y: 50}
var startPosition = model.Vec{X: 350, Y: 300}

//MaxFlightTime designates how long a simulation will run before restarting
var MaxFlightTime = 1000
var char = m.CellCharacter{
	NeuronBreeder: mutation.BuildNormalBreeder(
		0.00625, mutation.NormalyDistributedGeneCreator),
	Activator:   mutation.LogisticSigmoidActivator,
	GeneCreator: mutation.NormalyDistributedGeneCreator,
}

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
