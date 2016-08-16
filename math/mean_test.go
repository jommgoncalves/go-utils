package math

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestMean(t *testing.T) {
	inputsSum := []int{0, 100, 100}
	intputsCount := []int{0, 14, 17}
	expected := []int{0, 7, 6}

	for index, inputSum := range inputsSum {
		m := NewMean(inputSum, intputsCount[index])
		assert.Equal(t,
			m.Calc(),
			expected[index],
			fmt.Errorf("The mean should be equals to %d.", expected[index]))
	}
}
