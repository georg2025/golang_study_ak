package main

import "fmt"

// this is our base struct
type BaseMover struct {
	speed int
}

// and thats are his 2 childs
type FastMover struct {
	BaseMover
}

type SlowMover struct {
	BaseMover
}

// Our interface
type Mover interface {
	Move() string
	Speed() int
	MaxSpeed() int
	MinSpeed() int
}

// mostly we need to implement methods only on base struct.
// Exept for Move func that differ for fast and slow movers
func (mover FastMover) Move() string {
	return fmt.Sprintf("Fast mover! Moving at speed: %d", mover.speed)
}

func (mover SlowMover) Move() string {
	return fmt.Sprintf("Slow mover! Moving at speed: %d", mover.speed)
}

func (mover BaseMover) Speed() int {
	return mover.speed
}

func (mover BaseMover) MaxSpeed() int {
	return 120
}

func (mover BaseMover) MinSpeed() int {
	return 10
}

func main() {
	var movers []Mover
	fm := FastMover{BaseMover{100}}
	sm := SlowMover{BaseMover{10}}
	movers = append(movers, fm, sm)

	for _, mover := range movers {
		fmt.Println(mover.Move())
		fmt.Println("Maximum speed:", mover.MaxSpeed())
		fmt.Println("Minimum speed:", mover.MinSpeed())
	}
}
