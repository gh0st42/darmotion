package main

import (
	"fmt"
	"strconv"
)

type SWIM struct {
	radius                float32
	cell_distance         float32
	node_speed            float32
	wait_time_exp         float32
	wait_time_upper_bound float32
}

func onGenerateSWIM() {
	fmt.Println("Generating...")

	radius := strconv.FormatFloat(float64(swim.radius), 'f', -1, 32)
	cell_distance := strconv.FormatFloat(float64(swim.cell_distance), 'f', -1, 32)
	node_speed := strconv.FormatFloat(float64(swim.node_speed), 'f', -1, 32)
	wait_time_exp := strconv.FormatFloat(float64(swim.wait_time_exp), 'f', -1, 32)
	wait_time_upper_bound := strconv.FormatFloat(float64(swim.wait_time_upper_bound), 'f', -1, 32)

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
			base_cmd := baseCommand("SWIM", fout_i)
			run_cmd_args := fmt.Sprintf("%v -r %v -c %v -m %v -e %v -u %v", base_cmd, radius, cell_distance, node_speed, wait_time_exp, wait_time_upper_bound)
			execCommand(run_cmd_args)
			convert(fout_i)
		}
	} else {
		base_cmd := baseCommand("SWIM", fout)
		run_cmd_args := fmt.Sprintf("%v -r %v -c %v -m %v -e %v -u %v", base_cmd, radius, cell_distance, node_speed, wait_time_exp, wait_time_upper_bound)
		execCommand(run_cmd_args)
		convert(fout)
	}
}
