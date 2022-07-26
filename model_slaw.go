package main

import (
	"fmt"
	"strconv"
)

type SLAW struct {
	num_waypoints    int32
	min_pause        int32
	max_pause        int32
	levy_exp         float32
	hurst            float32
	dist_weight      float32
	clustering_range int32
	cluster_ratio    float32
	waypoint_ratio   float32
}

func onGenerateSLAW() {
	fmt.Println("Generating...")

	num_waypoints := strconv.Itoa(int(slaw.num_waypoints))
	min_pause := strconv.Itoa(int(slaw.min_pause))
	max_pause := strconv.Itoa(int(slaw.max_pause))
	levy_exp := strconv.FormatFloat(float64(slaw.levy_exp), 'f', -1, 32)
	hurst := strconv.FormatFloat(float64(slaw.hurst), 'f', -1, 32)
	dist_weight := strconv.Itoa(int(slaw.dist_weight))
	clustering_range := strconv.Itoa(int(slaw.clustering_range))
	cluster_ratio := strconv.FormatFloat(float64(slaw.cluster_ratio), 'f', -1, 32)
	waypoint_ratio := strconv.FormatFloat(float64(slaw.waypoint_ratio), 'f', -1, 32)

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
			base_cmd := baseCommand("SLAW", fout_i)
			run_cmd_args := fmt.Sprintf("%v -w %v -p %v -P %v -b %v -h %v -l %v -r %v -Q %v -W %v", base_cmd, num_waypoints, min_pause, max_pause, levy_exp, hurst, dist_weight, clustering_range, cluster_ratio, waypoint_ratio)
			execCommand(run_cmd_args)
			convert(fout_i)
		}
	} else {
		base_cmd := baseCommand("SLAW", fout)
		run_cmd_args := fmt.Sprintf("%v -w %v -p %v -P %v -b %v -h %v -l %v -r %v -Q %v -W %v", base_cmd, num_waypoints, min_pause, max_pause, levy_exp, hurst, dist_weight, clustering_range, cluster_ratio, waypoint_ratio)
		execCommand(run_cmd_args)
		convert(fout)
	}
}
