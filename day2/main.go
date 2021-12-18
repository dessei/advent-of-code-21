package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type command_num struct {
	command string
	num_num int
}
type position struct {
	hori  int
	verti int
	aim   int
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Reading File didn't work")
		return
	}
	lines := strings.Split(string(input), "\n")
	var cmd_nm []command_num
	cmd_nm = make([]command_num, 0, len(lines))
	for _, j := range lines {
		if len(j) == 0 {
			continue
		}
		tmp := strings.Split(j, " ")
		var tmp_num command_num
		tmp_num.command = tmp[0]
		tmp_num.num_num, _ = strconv.Atoi(tmp[1])
		cmd_nm = append(cmd_nm, tmp_num)
	}
	var pos position
	pos.hori = 0
	pos.verti = 0
	pos.aim = 0
	for _, j := range cmd_nm {
		switch j.command {
		case "forward":
			pos.hori += j.num_num
			pos.verti += j.num_num * pos.aim
		case "down":
			pos.aim += j.num_num
		case "up":
			pos.aim -= j.num_num
		}
	}
	res := pos.hori * pos.verti
	fmt.Println(res)
}

/*
type command_num struct {
	command string
	num_num int
}
type position struct {
	hori  int
	verti int
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Reading File didn't work")
		return
	}
	lines := strings.Split(string(input), "\n")
	var cmd_nm []command_num
	cmd_nm = make([]command_num, 0, len(lines))
	for _, j := range lines {
		if len(j) == 0 {
			continue
		}
		fmt.Println(j)
		tmp := strings.Split(j, " ")
		var tmp_num command_num
		tmp_num.command = tmp[0]
		tmp_num.num_num, _ = strconv.Atoi(tmp[1])
		cmd_nm = append(cmd_nm, tmp_num)
	}
	fmt.Println(cmd_nm)
	var pos position
	pos.hori = 0
	pos.verti = 0
	for _, j := range cmd_nm {
		switch j.command {
		case "forward":
			pos.hori += j.num_num
		case "down":
			pos.verti += j.num_num
		case "up":
			pos.verti -= j.num_num
		}
	}
	res := pos.hori * pos.verti
	fmt.Println(res)
}
*/
