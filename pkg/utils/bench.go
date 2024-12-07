package utils

import "time"

type Benchmark struct {
	start time.Time
	end   time.Time
}

func NewBenchmark() *Benchmark {
	return &Benchmark{start: time.Now(), end: time.Now()}
}

func (m *Benchmark) Start() {
	m.start = time.Now()
}

func (m *Benchmark) End() {
	m.end = time.Now()
}

func (m *Benchmark) Duration() time.Duration {
	return m.end.Sub(m.start)
}
