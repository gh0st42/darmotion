package main

import (
	"fmt"
	"strconv"
)

type RWP struct {
	max_speed int32
	min_speed int32
	max_pause int32
}

func onGenerateRandomWaypoint() {
	fmt.Println("Generating...")

	max_speed := strconv.Itoa(int(rwp.max_speed))
	min_speed := strconv.Itoa(int(rwp.min_speed))
	max_pause := strconv.Itoa(int(rwp.max_pause))
	var fout string
	if filename == "" {
		fout = "move"
	} else {
		fout = filename
	}
	if batch && reps > 0 && filename != "" {
		fmt.Println("Batch mode")
		for i := 0; i < int(reps); i++ {
			fout_i := fmt.Sprintf("%v-%v", fout, i+1)
			base_cmd := baseCommand("RandomWaypoint", fout_i)
			run_cmd_args := fmt.Sprintf("%v -h %v -l %v -p %v", base_cmd, max_speed, min_speed, max_pause)
			execCommand(run_cmd_args)
			convert(fout_i)
		}
	} else {
		base_cmd := baseCommand("RandomWaypoint", fout)
		run_cmd_args := fmt.Sprintf("%v -h %v -l %v -p %v", base_cmd, max_speed, min_speed, max_pause)
		execCommand(run_cmd_args)
		convert(fout)
	}
}
