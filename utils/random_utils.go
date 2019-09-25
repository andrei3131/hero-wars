package utils

import (
	"math/rand"
)

// Uniform int generator from range
type IntRange struct {
	min, max  int
	generator *rand.Rand
}

func NewIntRange(min, max int, generator *rand.Rand) *IntRange {
	return &IntRange{
		min:       min,
		max:       max,
		generator: generator,
	}
}

func (ir *IntRange) NextRandomInt() int {
	return ir.generator.Intn(ir.max-ir.min+1) + ir.min
}

// Uniform float32 generator from range
type FloatRange struct {
	min, max  float32
	generator *rand.Rand
}

func NewFloat32Range(min, max float32, generator *rand.Rand) *FloatRange {
	return &FloatRange{
		min:       min,
		max:       max,
		generator: generator,
	}
}

func (ir *FloatRange) NextRandomFloat32() float32 {
	return ir.min + ir.generator.Float32()*(ir.max-ir.min)
}
