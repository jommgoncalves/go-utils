package math

import (
	"fmt"
	"testing"

	"github.com/bmizerany/assert"
)

func TestRound(t *testing.T) {
	inputs := []float64{0, 1.5, 2.1, 1.55, -1.5}
	expected := []int{0, 2, 2, 2, -2}
	for index, input := range inputs {
		assert.Equal(t,
			Round(input),
			expected[index],
			fmt.Errorf("The mean should be equals to %d.", expected[index]))
	}
}
