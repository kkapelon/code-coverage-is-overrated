package velocity

import (
	"fmt"
	"testing"
)

func TestVelocityScenarios(t *testing.T) {

	parameters := []struct{ exampleAngle, exampleDirection, expectedVelocity int }{
		{3, 4, 5},
		{10, 2, 2},
		{8, 3, 2},
		{6, 6, 3},
		{12, 7, 3},
		{22, 8, 2},
		{1, 20, 1},
	}

	for i := range parameters {
		t.Run(fmt.Sprintf("Testing [%v]", i), func(t *testing.T) {
			actual := calculateVelocity(parameters[i].exampleAngle, parameters[i].exampleDirection)
			if actual != parameters[i].expectedVelocity {
				t.Logf("expected%d: , actual:%d", parameters[i].expectedVelocity, actual)
				t.Fail()
			}
		})
	}

}
