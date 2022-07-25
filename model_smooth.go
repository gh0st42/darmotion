package main

import (
	"fmt"
	"strconv"
)

type SMOOTH struct {
	comm_range int32
	clusters   int32
	alpha      float32
	beta       float32
	f_min      int32
	f_max      int32
	p_min      int32
	p_max      int32
}

func onGenerateSMOOTH() {
	fmt.Println("Generating...")
	comm_range := strconv.Itoa(int(smooth.comm_range))
	clusters := strconv.Itoa(int(smooth.clusters))
	alpha := strconv.FormatFloat(float64(smooth.alpha), 'f', -1, 32)
	beta := strconv.FormatFloat(float64(smooth.beta), 'f', -1, 32)
	f_min := strconv.Itoa(int(smooth.f_min))
	f_max := strconv.Itoa(int(smooth.f_max))
	p_min := strconv.Itoa(int(smooth.p_min))
	p_max := strconv.Itoa(int(smooth.p_max))

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
			base_cmd := baseCommand("SMOOTH", fout_i)
			run_cmd_args := fmt.Sprintf("%v-g %v -h %v -k %v -l %v -m %v -o %v -p %v -q %v", base_cmd, comm_range, clusters, alpha, f_min, f_max, beta, p_min, p_max)
			execCommand(run_cmd_args)
			convert(fout_i)
		}
	} else {
		base_cmd := baseCommand("SMOOTH", fout)
		run_cmd_args := fmt.Sprintf("%v -g %v -h %v -k %v -l %v -m %v -o %v -p %v -q %v", base_cmd, comm_range, clusters, alpha, f_min, f_max, beta, p_min, p_max)
		execCommand(run_cmd_args)
		convert(fout)
	}
}
