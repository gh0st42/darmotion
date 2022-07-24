package main

import (
	"fmt"
	"os/exec"
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
	d := strconv.Itoa(int(duration))
	s := strconv.Itoa(int(skip))
	n := strconv.Itoa(int(num_nodes))
	w := strconv.Itoa(int(width))
	h := strconv.Itoa(int(height))
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
			cmd, err := exec.Command(BONNMOTION, "-f", fout_i, "SMOOTH", "-i", s, "-d", d, "-n", n, "-x", w, "-y", h, "-g", comm_range, "-h", clusters, "-k", alpha, "-l", f_min, "-m", f_max, "-o", beta, "-p", p_min, "-q", p_max).Output()
			fmt.Println(string(cmd))
			if err != nil {
				fmt.Println(err)
				return
			}
			convert(fout_i)
		}
	} else {
		cmd, err := exec.Command(BONNMOTION, "-f", fout, "SMOOTH", "-i", s, "-d", d, "-n", n, "-x", w, "-y", h, "-g", comm_range, "-h", clusters, "-k", alpha, "-l", f_min, "-m", f_max, "-o", beta, "-p", p_min, "-q", p_max).Output()
		fmt.Println(string(cmd))
		if err != nil {
			fmt.Println(err)
			return
		}
		convert(fout)
	}
}
