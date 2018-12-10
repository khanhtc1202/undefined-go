package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

// MutationRate is the rate of mutation
var MutationRate = 0.0004

// PopSize is the size of the population
var PopSize = 250

// PoolSize is the max size of the pool
var PoolSize = 30

// FitnessLimit is the fitness of the evolved image we are satisfied with
var FitnessLimit int64 = 7500

type Organism struct {
	DNA     []byte
	Fitness float64
}

func (o *Organism) calcFitness(target []byte) {
	score := 0
	for i := 0; i < len(o.DNA); i++ {
		if o.DNA[i] == target[i] {
			score++
		}
	}
	o.Fitness = float64(score) / float64(len(o.DNA))
}

func (o *Organism) mutate() {
	for i := 0; i < len(o.DNA); i++ {
		if rand.Float64() < MutationRate {
			o.DNA[i] = byte(rand.Intn(95) + 32)
		}
	}
}

func createOrganism(target []byte) (organism Organism) {
	ba := make([]byte, len(target))
	for i := 0; i < len(target); i++ {
		ba[i] = byte(rand.Intn(95) + 32)
	}
	organism = Organism{
		DNA:     ba,
		Fitness: 0,
	}
	organism.calcFitness(target)
	return
}

func createPopulation(target []byte) (population []Organism) {
	population = make([]Organism, PopSize)
	for i := 0; i < len(target); i++ {
		population[i] = createOrganism(target)
	}
	return
}

func createPool(population []Organism, target []byte, maxFitness float64) (pool []Organism) {
	pool = make([]Organism, 0)
	for i := 0; i < len(population); i++ {
		population[i].calcFitness(target)
		num := int((population[i].Fitness / maxFitness) * 100)
		for n := 0; n < num; n++ {
			pool = append(pool, population[i])
		}
	}
	return
}

func crossover(d1 Organism, d2 Organism) Organism {
	child := Organism{
		DNA:     make([]byte, len(d1.DNA)),
		Fitness: 0,
	}
	mid := rand.Intn(len(d1.DNA))
	for i := 0; i < len(d1.DNA); i++ {
		if i > mid {
			child.DNA[i] = d1.DNA[i]
		} else {
			child.DNA[i] = d2.DNA[i]
		}

	}
	return child
}

func naturalSelection(pool []Organism, population []Organism, target []byte) []Organism {
	next := make([]Organism, len(population))

	for i := 0; i < len(population); i++ {
		r1, r2 := rand.Intn(len(pool)), rand.Intn(len(pool))
		a := pool[r1]
		b := pool[r2]

		child := crossover(a, b)
		child.mutate()
		child.calcFitness(target)

		next[i] = child
	}

	return next
}

func getBest(population []Organism) Organism {
	best := float64(0)
	index := 0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > best {
			index = i
			best = population[i].Fitness
		}
	}
	return population[index]
}

func main() {
	start := time.Now()
	rand.Seed(time.Now().UTC().UnixNano())

	target := []byte("To be or not to be")
	population := createPopulation(target)

	found := false
	generation := 0
	for !found {
		generation++
		bestOrganism := getBest(population)
		fmt.Printf("\r generation: %d | %s | fitness: %2f", generation, string(bestOrganism.DNA), bestOrganism.Fitness)

		if bytes.Compare(bestOrganism.DNA, target) == 0 {
			found = true
		} else {
			maxFitness := bestOrganism.Fitness
			pool := createPool(population, target, maxFitness)
			population = naturalSelection(pool, population, target)
		}

	}
	elapsed := time.Since(start)
	fmt.Printf("\nTime taken: %s\n", elapsed)
}
