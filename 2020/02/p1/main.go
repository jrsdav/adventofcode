package main

import (
	"fmt"
	"strconv"
	"strings"
)

type password struct {
	min, max         int
	letter, password string
}

var correct = 0
var incorrect = 0

func main() {

	input := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`

	items := strings.Split(input, "\n")

	passwords := []*password{}

	for _, v := range items {
		fields := strings.Fields(v)

		minmax := strings.Split(fields[0], "-")
		letter := strings.Split(fields[1], ":")

		min, _ := strconv.Atoi(minmax[0])
		max, _ := strconv.Atoi(minmax[1])

		passwords = append(passwords, &password{
			min:      min,
			max:      max,
			letter:   string(letter[0]),
			password: string(fields[2]),
		})

	}
	for _, pw := range passwords {
		fmt.Printf("Checking password: %s...\nRequires: '%s', '%v'-'%v' times\n", pw.password, pw.letter, pw.min, pw.max)
		if count := strings.Count(pw.password, pw.letter); count >= pw.min && count <= pw.max {
			fmt.Printf("Password %s is correct!\n\n", pw.password)
			correct++
		} else {
			fmt.Printf("Password: %s does not meet requirements!\n\n", pw.password)
			incorrect++
		}
	}
	fmt.Printf("Found %v correct and %v incorrect passwords\n", correct, incorrect)
}
