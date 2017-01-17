package rocket

import (
	"github.com/JamesRCooper/NeuralRockets/model"
	neural "github.com/JamesRCooper/geneticNeuralNetwork/model"
)

//Rocket can fly!!!
type Rocket struct {
	position     model.Vec
	velocity     model.Vec
	acceleration model.Vec

	Brain neural.Network

	boundryCheck CheckBoundry
	HitBoundry   bool
	FlightTime   int

	goalPosition *model.Vec
}

//BuildMapping builds out the positional parameter for display
func (r *Rocket) BuildMapping() model.Pos {
	return model.Pos{X: r.position.X, Y: r.position.Y, A: r.velocity.Angle()}
}

//Update updates the position, velocity, and acceleration of a rocket
func (r *Rocket) Update() error {
	if r.HitBoundry {
		return nil
	}
	if !r.boundryCheck(r.position) {
		r.HitBoundry = true
		return nil
	}

	defer func(r *Rocket) { r.FlightTime++ }(r)
	inputs := buildInputs(r)

	output, err := r.Brain.CalculateOutput(inputs)
	if err != nil {
		return err
	}

	timePerStep := 1.0
	deltaPos := r.velocity.Copy()
	deltaPos.Mult(timePerStep)
	r.position.Add(deltaPos)
	r.velocity.Add(model.Vec{X: output[0] * timePerStep, Y: output[1] * timePerStep})

	return nil
}

func buildInputs(r *Rocket) []float64 {

	inputs := make([]float64, 10)

	midPos := r.position
	inputs[0] = -1.0 * midPos.X
	inputs[1] = -1.0 * midPos.Y

	deltaPos := 5.0
	inputs[2] = openSpace(r, model.Vec{X: midPos.X - deltaPos, Y: midPos.Y}) //up
	inputs[3] = openSpace(r, model.Vec{X: midPos.X + deltaPos, Y: midPos.Y}) //down
	inputs[4] = openSpace(r, model.Vec{X: midPos.X, Y: midPos.Y - deltaPos}) //left
	inputs[5] = openSpace(r, model.Vec{X: midPos.X, Y: midPos.Y + deltaPos}) //right

	inputs[6] = -1.0 * r.velocity.X
	inputs[7] = -1.0 * r.velocity.Y

	inputs[8] = -1.0 * r.goalPosition.X
	inputs[9] = -1.0 * r.goalPosition.Y

	return inputs
}

func openSpace(r *Rocket, v model.Vec) float64 {
	if r.boundryCheck(v) {
		return 0.0
	}
	return -1.0
}

//CheckBoundry is a functional interface for a function that determines whether
//a point exists in an open or restricted region
type CheckBoundry func(model.Vec) bool

func (r *Rocket) breed(partnerRocket Rocket) {
	r.Brain.Breed(&partnerRocket.Brain)
}
