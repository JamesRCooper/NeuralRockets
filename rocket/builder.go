package rocket

import (
	neural "github.com/jamescooper/neural/model"
	"github.com/jamescooper/neuralRockets/model"
)

//Builder holds all information for creating a population of rockets, and has
//methods for creating those populations
type Builder struct {
	numberOfInputs  int
	layerSizes      []int
	characteristics *neural.CellCharacter

	goalPosition  *model.Vec
	startPosition model.Vec

	boundryChecker CheckBoundry
}

//CreateBuilder constructs a new builder for creating a population
func CreateBuilder(
	numberOfInputs int,
	layerSizes []int,
	char *neural.CellCharacter,
	goalPosition *model.Vec,
	startPosition model.Vec,
	boundryChecker CheckBoundry) (builder Builder) {

	//builder := new(Builder)
	builder.numberOfInputs = numberOfInputs
	builder.layerSizes = layerSizes
	builder.characteristics = char
	builder.goalPosition = goalPosition
	builder.startPosition = startPosition
	builder.boundryChecker = boundryChecker
	return
}

//Build creates a new rocket given a boundry checking algorithm and set of
//characteristics
func (b Builder) Build() *Rocket {
	rocket := new(Rocket)
	rocket.goalPosition = b.goalPosition
	rocket.position = b.startPosition
	rocket.velocity = model.Vec{X: 0, Y: 0}
	rocket.acceleration = model.Vec{X: 0, Y: 0}
	rocket.HitBoundry = false

	rocket.Brain = neural.NewNetwork(b.numberOfInputs, b.layerSizes, b.characteristics)

	rocket.boundryCheck = b.boundryChecker

	return rocket
}
