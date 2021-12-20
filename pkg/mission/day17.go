package mission

type Day17 struct {
	input []string
	x1    int
	x2    int
	y1    int
	y2    int
}

func (d *Day17) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day17) Task1() int {
	return (0 - d.y1 - 1) * (0 - d.y1) / 2
}

func (d *Day17) Task2() int {
	count := 0
	found := make(map[Coord]bool)
	for targetX := d.x1; targetX <= d.x2; targetX++ {
		for targetY := d.y1; targetY <= d.y2; targetY++ {
			for startingVelocityX := 1; startingVelocityX <= targetX; startingVelocityX++ {
				for startingVelocityY := d.y1; startingVelocityY <= -d.y1-1; startingVelocityY++ {
					trajectoryX := 0
					trajectoryY := 0
					for i := 0; trajectoryY > targetY && trajectoryX <= targetX; i++ {
						if startingVelocityX-i > 0 {
							trajectoryX += startingVelocityX - i
						}
						trajectoryY += startingVelocityY - i
						if trajectoryX == targetX && trajectoryY == targetY && !found[Coord{startingVelocityX, startingVelocityY}] {
							count++
							found[Coord{startingVelocityX, startingVelocityY}] = true
							break
						}
					}
				}
			}
		}
	}
	return count
}

func (d *Day17) parseInput() {
	d.x1 = 155
	d.x2 = 182
	d.y1 = -117
	d.y2 = -67
}
