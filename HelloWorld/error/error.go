package main

import (
	"fmt"
	"time"
)

type customError struct {
	errCode string
	errMsg string
	timestamp time.Time
}

func (e customError) Error() string {
	return fmt.Sprintf("[%v] - ErrorCode: %s, ErrorMsg: %s.", e.timestamp.Nanosecond(), e.errCode, e.errMsg)
}

func divide(numerator float64, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, customError {
			errCode: "DIVIDEBYZERO",
			errMsg: "Division by zero is not allowed",
			timestamp: time.Now(),
		}
	}
	return numerator/denominator, nil
}

func main() {
	// Defining an slice of maps
	// numPairs := []map[string]float64 {
	// 	{"numerator": 10, "denominator": 2},
	// 	{"numerator": 12, "denominator": 0},
	// 	{"numerator": 12, "denominator": 3},
	// }
	
	// Aleternative - better to use struct so that keys are same.
	type numStruct struct {
		numerator float64
		denominator float64
	}
	numPairs := []numStruct {
		{10, 2},
		numStruct{12, 0},
		numStruct{denominator: 0, numerator: 10},
		{numerator: 12, denominator: 4},
	}

	for _, m := range numPairs {
		// Use below if you are using slice of maps.
		//r, err := divide(m["numerator"], m["denominator"])

		// Use below if you are using slice of struct.
		r, err := divide(m.numerator, m.denominator)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%f/%f = %f\n",m.numerator, m.denominator, r)
		}
	}
}