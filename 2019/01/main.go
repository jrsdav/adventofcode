// Run in https://play.golang.org/
// https://adventofcode.com/2019/day/1
// Contains solutions for part 1 and part 2
package main

import (
	"fmt"
	"math"
)

var part1 float64
var part2 float64

func main() {
	mass := []float64{58444, 100562, 133484, 67910, 58372, 104607, 108786, 137410, 62910, 76115, 64142, 59324, 54327, 92864, 94120, 63931, 128696, 111758, 65698, 54930, 116136, 127111, 133914, 52992, 90364, 107637, 62118, 147901, 62347, 53614, 140690, 115587, 66148, 95729, 148847, 84269, 71569, 85026, 130871, 102470, 53328, 63308, 104085, 57744, 123008, 120983, 94968, 69402, 83830, 137069, 121062, 71267, 103035, 97604, 129153, 65595, 148655, 124573, 139257, 59722, 101050, 139557, 74362, 50024, 101750, 83209, 117840, 139442, 127810, 113438, 94731, 125471, 96653, 88522, 125573, 74456, 89839, 84458, 128451, 68608, 92504, 103549, 117980, 126850, 144675, 59752, 60986, 125867, 89982, 108717, 134815, 89209, 143434, 61123, 103162, 139529, 122228, 137866, 78676, 80779}

	for _, v := range mass {
		fuel := fuelCounterUpperFormula(v)
		part1 += fuel
	}
	fmt.Printf("Part 1: %d\n", int(part1))

	for _, v := range mass {
		fuel := fuelCounterRecursive(v)
		part2 += fuel
	}
	fmt.Printf("Part 2: %d\n", int(part2))
}

func fuelCounterUpperFormula(n float64) float64 {
	fuel := math.Floor(float64(n/3)) - 2
	return fuel
}

func fuelCounterRecursive(n float64) float64 {
	newFuel := fuelCounterUpperFormula(n)

	if newFuel <= 0 {
		return 0
	} else {
		newFuel += fuelCounterRecursive(newFuel)
	}
	return newFuel
}
