package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This function generates an event based on a probability
// then returns the consequence of the event if it occurs
func eventGen(probability float64, consequence float64) float64 {
	rand.Seed(time.Now().UnixNano())
	var guess int = rand.Intn(100)
	var probThreshold int = int(probability * 100)
	if guess <= probThreshold {
		return consequence
	} else {
		return 0
	}
}

// sum array numbers
func arraySum(arr []float64) float64 {
	var result float64 = 0
	for i := 0; i < len(arr); i++ {
		result += arr[i]
	}
	return result
}

// struct stores risk information: probability and consequence used in the fxn ^
type risk struct {
	prob float64
	cons float64
}

func main() {
	// init risk array of structs TODO: read in a csv to this
	risks := []risk{
		{
			prob: 0.67,
			cons: 1,
		},
		{
			prob: 0.67,
			cons: 2,
		},
		{
			prob: 0.67,
			cons: 2,
		},
		{
			prob: 0.67,
			cons: 2,
		},
		{
			prob: 0.67,
			cons: 1,
		},
		{
			prob: 0.67,
			cons: 2,
		},
		{
			prob: 0.67,
			cons: 1,
		},
		{
			prob: 0.67,
			cons: 1,
		},
		{
			prob: 0.67,
			cons: 2,
		},
		{
			prob: 0.67,
			cons: 2,
		},
	}

	totalRuns := 100000
	var runResults []float64
	var runSums []float64
	var threshold int = int(0.25 * float64(len(risks)) * 5)

	for j := 0; j < totalRuns; j++ {
		// declare an empty array of floats to store events from a single run
		var events []float64

		// run through each risk and record each event
		for i := 0; i < len(risks); i++ {
			var event float64 = eventGen(risks[i].prob, risks[i].cons)
			events = append(events, event)
		}
		var runSum float64 = arraySum(events)
		runSums = append(runSums, runSum)
		if int(runSum) >= threshold {
			runResults = append(runResults, 1)
		} else {
			runResults = append(runResults, 0)
		}
	}

	fmt.Println("Total Runs: ", totalRuns,
		"Failures: ", arraySum(runResults),
		"Failure Rate: ", arraySum(runResults)/float64(len(runSums)))
}
