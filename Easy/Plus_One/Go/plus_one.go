package main

//----- REQUIRED IMPORTS -----\\
import (
	"math/big"
	"strconv"
)

func main() {
	// ----- ONLY HERE FOR TESTING ----- \\
}

// ========== THE SOLUTION FUNCTION ========== \\
func plusOne(digits []int) []int {
	sum := ""
	for _, v := range digits {
		sum += strconv.Itoa(v)
	}

	bigSum := new(big.Int)
	bigSum.SetString(sum, 10)

	bigSum.Add(bigSum, big.NewInt(1))

	sumStr := bigSum.String()

	output := []int{}
	for _, ch := range sumStr {
		num, _ := strconv.Atoi(string(ch))
		output = append(output, num)
	}

	return output
}
