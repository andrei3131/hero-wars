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

// Uniform float64 generator from range
type FloatRange struct {
	min, max  float64
	generator *rand.Rand
}

func NewFloat64Range(min, max float64, generator *rand.Rand) *FloatRange {
	return &FloatRange{
		min:       min,
		max:       max,
		generator: generator,
	}
}

func (ir *FloatRange) NextRandomFloat64() float64 {
	return ir.min + ir.generator.Float64()*(ir.max-ir.min)
}
