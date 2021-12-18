package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func part_1() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Reading input didn't work!")
		return
	}
	lines := strings.Split(string(content), "\n")
	var LineLength int = len(lines[0])
	var OneIncidence []int = make([]int, LineLength)
	for _, j := range lines {
		for i := 0; i < LineLength; i++ {
			if j[i] == '1' {
				OneIncidence[i] += 1
			}
		}
	}
	gamma := 0
	epsilon := 0
	for i := 0; i < LineLength; i++ {
		if OneIncidence[i] > (len(lines) / 2) {
			gamma += int(math.Pow(2, float64(LineLength-i-1)))
		} else {
			epsilon += int(math.Pow(2, float64(LineLength-i-1)))
		}
	}
	fmt.Println(gamma * epsilon)
}

func main() {
	part_1()
	fmt.Println("OwO world")
}
