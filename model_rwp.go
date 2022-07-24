package main

import (
	"fmt"
	"os/exec"
	"strconv"
)

func onGenerateRandomWaypoint() {
	fmt.Println("Generating...")
	// -f smooth SMOOTH -d 43200 -n 100 -x 8000 -y 8000 -g 10 -h 40 -k 1.45 -l 1 -m 14000 -o 1.5 -p 10 -q 360
	d := strconv.Itoa(int(duration))
	n := strconv.Itoa(int(num_nodes))
	w := strconv.Itoa(int(width))
	h := strconv.Itoa(int(height))
	s := strconv.Itoa(int(skip))

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
			cmd, err := exec.Command(BONNMOTION, "-f", fout_i, "RandomWaypoint", "-i", s, "-d", d, "-n", n, "-x", w, "-y", h).Output()
			fmt.Println(string(cmd))
			if err != nil {
				fmt.Println(err)
				return
			}
			convert(fout_i)
		}
	} else {
		cmd, err := exec.Command(BONNMOTION, "-f", fout, "RandomWaypoint", "-i", s, "-d", d, "-n", n, "-x", w, "-y", h).Output()
		fmt.Println(string(cmd))
		if err != nil {
			fmt.Println(err)
			return
		}
		convert(fout)
	}
}
