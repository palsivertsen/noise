package noise

import (
	"log"
	"math"
	"math/rand"
)

// Smooth is a semi random number generator.
type Smooth struct {
	// Size of numbers to pre generate.
	// The higher this number, the smoother values returned by Next()
	// New values are automatically generated when values are exhausted.
	Size int

	// Optionally set a rand to use
	Rand *rand.Rand

	r    *rand.Rand
	pool []float64
}

// Next returns the next number
func (s *Smooth) Next() float64 {
	if s.Size == 0 {
		log.Panicf("Size (%d) needs to be larger than 0", s.Size)
	}

	if len(s.pool) <= 1 {
		s.gen()
	}

	r := s.pool[0]
	s.pool = s.pool[1:]

	return r
}

func (s *Smooth) gen() {
	var start float64
	if len(s.pool) == 0 {
		start = s.rand().Float64()
	} else {
		start = s.pool[0]
	}

	if s.Size > cap(s.pool) {
		s.pool = make([]float64, s.Size)
	} else {
		s.pool = s.pool[:s.Size]
	}

	s.pool[0] = start
	s.pool[len(s.pool)-1] = s.rand().Float64()

	s.fill(s.pool)
}

func (s *Smooth) fill(p []float64) {
	if len(p) < 3 {
		return
	}

	midI := len(p) / 2

	randF := 1 / float64(s.Size) * float64(len(p))

	p[midI] = s.randRange(randF, p[0], p[len(p)-1])

	s.fill(p[:midI+1])
	s.fill(p[midI:])
}

func (s *Smooth) randRange(randF, a, b float64) float64 {
	min := math.Min(a, b)
	max := math.Max(a, b)

	min = min - (randF * min)
	if min < 0 {
		min = 0
	}
	max = max + (randF * max)
	if max > 1 {
		max = 1
	}

	r := s.rand().Float64()
	return r*(max-min) + min
}

func (s *Smooth) rand() *rand.Rand {
	if s.Rand != nil {
		return s.Rand
	}
	if s.r == nil {
		s.r = rand.New(rand.NewSource(1))
	}
	return s.r
}
