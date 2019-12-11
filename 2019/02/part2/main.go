// Run in https://play.golang.org/
// https://adventofcode.com/2019/day/2

package main

import (
	"fmt"
	"log"

	"github.com/bradfitz/iter"
)

func main() {
	for i := range iter.N(99) {
		for n := range iter.N(99) {
			input := IntCode()
			input[1] = i
			input[2] = n
			result, err := IntCodeProgram(input)
			if err != nil {
				log.Fatal(err)
			}
			if result[0] == 19690720 {
				fmt.Printf("Noun: %d\nVerb: %d\n", result[1], result[2])
				nv := 100 * result[1]
				newr := nv + result[2]
				fmt.Printf("Computed: %d\n", newr)
			}
		}
	}
}

func IntCode() []int {
	return []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 6, 1, 19, 1, 5, 19, 23, 2, 6, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 5, 35, 39, 2, 6, 39, 43, 2, 6, 43, 47, 1, 5, 47, 51, 2, 9, 51, 55, 1, 5, 55, 59, 1, 10, 59, 63, 1, 63, 6, 67, 1, 9, 67, 71, 1, 71, 6, 75, 1, 75, 13, 79, 2, 79, 13, 83, 2, 9, 83, 87, 1, 87, 5, 91, 1, 9, 91, 95, 2, 10, 95, 99, 1, 5, 99, 103, 1, 103, 9, 107, 1, 13, 107, 111, 2, 111, 10, 115, 1, 115, 5, 119, 2, 13, 119, 123, 1, 9, 123, 127, 1, 5, 127, 131, 2, 131, 6, 135, 1, 135, 5, 139, 1, 139, 6, 143, 1, 143, 6, 147, 1, 2, 147, 151, 1, 151, 5, 0, 99, 2, 14, 0, 0}
}

func IntCodeProgram(i []int) ([]int, error) {
	position := 0

	for {
		opcode := i[position]

		if opcode == 99 {
			break
		}

		p1 := i[position+1]
		p2 := i[position+2]
		p3 := i[position+3]

		switch opcode {
		case 1:
			result := i[p1] + i[p2]
			i[p3] = result
		case 2:
			result := i[p1] * i[p2]
			i[p3] = result
		default:
			return nil, fmt.Errorf("opcode %d at position %d unrecognized", opcode, position)
		}
		position += 4
	}
	return i, nil
}
