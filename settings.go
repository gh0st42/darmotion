package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

const (
	BONNMOTION = "bonnmotion"
)

const (
	FORMAT_NSFILE = iota
	FORMAT_INTERVALFORMAT
	FORMAT_ONE
	FORMAT_BONNMOTION
)

var (
	multiline string
	filename  string
	selected  int32  = FORMAT_ONE
	formats          = []string{"NSFile", "IntervalFormat", "ONE", "Bonnmotion"}
	batch            = false
	reps      int32  = 1
	num_nodes int32  = 100
	skip      int32  = 0
	width     int32  = 8000
	height    int32  = 8000
	duration  int32  = 43200
	smooth    SMOOTH = SMOOTH{
		comm_range: 100,
		clusters:   40,
		alpha:      1.45,
		beta:       1.5,
		f_min:      1,
		f_max:      14000,
		p_min:      10,
		p_max:      3600,
	}
	rwp RWP = RWP{
		max_speed: 5,
		min_speed: 1,
		max_pause: 3600,
	}
	swim SWIM = SWIM{
		radius:                0.2,
		cell_distance:         0.5,
		node_speed:            1.5,
		wait_time_exp:         1.55,
		wait_time_upper_bound: 100,
	}
)

func baseCommand(model string, fout string) string {
	d := strconv.Itoa(int(duration))
	n := strconv.Itoa(int(num_nodes))
	w := strconv.Itoa(int(width))
	h := strconv.Itoa(int(height))
	s := strconv.Itoa(int(skip))
	return fmt.Sprintf("%v -f %v %v -d %v -i %v -n %v -x %v -y %v", BONNMOTION, fout, model, d, s, n, w, h)
}

func execCommand(run_cmd string) {
	fmt.Println("exec: ", run_cmd)
	args := strings.Fields(run_cmd)
	cmd, err := exec.Command(args[0], args[1:]...).Output()
	fmt.Println("result: ", string(cmd))
	if err != nil {
		fmt.Println(err)
		return
	}
}
