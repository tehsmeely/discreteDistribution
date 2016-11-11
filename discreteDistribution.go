package discreteDistribution
//http://keithschwarz.com/darts-dice-coins/

// func generate(probabilities, outcomes []float32) float32{
// 	if len(probabilities) != len(outcomes){
// 		//error
// 	}

// 	//check PDF is normalised to 100:
// 	if 
// }

import (
	"fmt"
	"math/rand"
)

func Generate(probabilities, outcomes []int) (int, bool) {
	var ok bool
	//terribly inefficient
	if len(probabilities) != len(outcomes){
		fmt.Println("slices not of same length")
		return -1, ok
	}

	//check PDF is normalised to 100:
	if sum(probabilities) != 100{
		fmt.Println("sum of probabilities != 100")
		return -1, ok
	}

	randomNum := rand.Intn(100)
	var runningTotal int
	returnVal := -1
	for i, val := range probabilities {
		if randomNum >= runningTotal && randomNum < runningTotal+val {
			returnVal = outcomes[i]
			ok = true
			break
		} else {
			runningTotal += val
		}
	}
	return returnVal, ok
}	

func sum(in []int) int {
	var sum int
	for _, val := range in {
		sum+=val
	}
	return sum
}