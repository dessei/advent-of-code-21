package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Reading File didn't work")
		return
	}
	lines := strings.Split(string(input), "\r\n")
	var nums []int
	nums = make([]int, 0, len(lines))
	for i, l := range lines {
		if len(l) == 0 {
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			fmt.Print(err)
			fmt.Printf("error at line %d \n", i)
			fmt.Println("converting number failed")
			return
		}
		nums = append(nums, n)
	}
	var increases int = 0
	var upper_check_value int = nums[0] + nums[1] + nums[2]
	var lower_check_value int
	for i := 3; i < len(nums); i++ {
		lower_check_value = nums[i] + nums[i-1] + nums[i-2]
		if lower_check_value > upper_check_value {
			increases = increases + 1
		}
		upper_check_value = lower_check_value
	}
	fmt.Println(increases)
}
