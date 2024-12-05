package main

import (
	"fmt"
)

// TODO: implement the function computeGrade

func main() {
	// TODO: call the function computeGrade
	fmt.Print(computeGrade(17.5, 28)) // 4.125
	fmt.Print(computeGrade(29, 28))   // Error
}

func computeGrade(achievedPoints float64, maxPoints float64) (float64, error) {
	if achievedPoints < 0 || maxPoints < 0 || achievedPoints > maxPoints {
		return 0.0, fmt.Errorf("IllegalArgumentException: Achieved Points (%v) or Max Points (%v) is invalid", achievedPoints, maxPoints)
	}
	return achievedPoints/maxPoints*5 + 1, nil
}
