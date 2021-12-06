package mission

import (
	"strconv"
	"strings"
)

type lanternFish struct {
	timer  int
	amount int
}

type fishTracker struct {
	FishByTimer map[int]*lanternFish
}

func (t *fishTracker) addFish(timer int, amount int) {
	if fishWithSameTimer, exists := t.FishByTimer[timer]; exists {
		fishWithSameTimer.amount += amount
	} else {
		t.FishByTimer[timer] = &lanternFish{timer: timer, amount: amount}
	}
}

func (t *fishTracker) getCount() int {
	totalAmount := 0
	for _, f := range t.FishByTimer {
		totalAmount += f.amount
	}
	return totalAmount
}

func newFishTracker() fishTracker {
	return fishTracker{FishByTimer: make(map[int]*lanternFish, 9)}
}

type Day6 struct {
	input []string
	fish  fishTracker
}

func (d *Day6) SetInput(input []string) {
	d.input = input
	d.parseInput()
}

func (d *Day6) Task1() int {
	d.simulatePopulation(80)
	return d.fish.getCount()
}

func (d *Day6) Task2() int {
	d.simulatePopulation(256 - 80)
	return d.fish.getCount()
}

func (d *Day6) simulatePopulation(days int) {
	for i := 0; i < days; i++ {
		newFish := newFishTracker()
		for _, currentFish := range d.fish.FishByTimer {
			if currentFish.timer > 0 {
				currentFish.timer--
			} else {
				currentFish.timer = 6
				newFish.addFish(8, currentFish.amount)
			}
			newFish.addFish(currentFish.timer, currentFish.amount)
		}
		d.fish = newFish
	}
}

func (d *Day6) parseInput() {
	strFish := strings.Split(d.input[0], ",")
	d.fish = newFishTracker()
	for _, f := range strFish {
		value, _ := strconv.Atoi(f)
		d.fish.addFish(value, 1)
	}
}
