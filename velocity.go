package velocity

func calculateVelocity(angle int, direction int) int {
	return ((3 * (4*angle - direction)) * 3) / (7 * (direction - (2 * angle))) * -1
}
