package velocity

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	angle := 5
	direction := 3
	velocity := calculateVelocity(angle, direction)
	if velocity != 4 {
		t.Fatalf(`Wanted %d, got %d`, 4, velocity)
	}
}
