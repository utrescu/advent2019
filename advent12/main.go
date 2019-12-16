package main

import "fmt"

func main() {
	fmt.Println("---------------------")
	positions := []*Coord{
		&Coord{0, 6, 1, 0, 0, 0},
		&Coord{4, 4, 19, 0, 0, 0},
		&Coord{-11, 1, 8, 0, 0, 0},
		&Coord{2, 19, 15, 0, 0, 0},
	}

	fmt.Printf("Energy: %d\n", part1(1000, positions))

}

// Coord stores position and velocity
type Coord struct {
	X  int
	Y  int
	Z  int
	VX int
	VY int
	VZ int
}

// Move moves the moon
func (c *Coord) Move() {
	c.X += c.VX
	c.Y += c.VY
	c.Z += c.VZ
}

// Print prints values
func (c *Coord) Print() {
	fmt.Printf("Position (%d,%d,%d) Velocity(%d,%d,%d)\n", c.X, c.Y, c.Z, c.VX, c.VY, c.VZ)
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}

func changeVelocity(one *Coord, two *Coord) {

	if one.X > two.X {
		one.VX += -1
		two.VX++
	} else if one.X < two.X {
		one.VX++
		two.VX += -1
	}

	if one.Y > two.Y {
		one.VY += -1
		two.VY++
	} else if one.Y < two.Y {
		one.VY++
		two.VY += -1
	}

	if one.Z > two.Z {
		one.VZ += -1
		two.VZ++
	} else if one.Z < two.Z {
		one.VZ++
		two.VZ += -1
	}

}

func part1(steps int, moons []*Coord) int {

	total := 0

	for step := 1; step <= steps; step++ {
		// Gravity
		// fmt.Printf("----- Step %d\n", step)
		for index := 0; index < len(moons); index++ {
			moon := moons[index]
			for index2 := index + 1; index2 < len(moons); index2++ {
				changeVelocity(moon, moons[index2])
			}
			// Velocity
			moon.Move()
			// moon.Print()
		}

		// fmt.Println()

	}

	for _, moon := range moons {
		pot := abs(moon.X) + abs(moon.Y) + abs(moon.Z)
		kin := abs(moon.VX) + abs(moon.VY) + abs(moon.VZ)
		total += pot * kin
	}

	return total
}
