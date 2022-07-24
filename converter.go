package main

import (
	"fmt"
	"os/exec"
)

func convert(filename string) {
	switch selected {
	case FORMAT_INTERVALFORMAT:
		convertToIF(filename)
	case FORMAT_ONE:
		convertToOne(filename)
	case FORMAT_NSFILE:
		convertToNSFile(filename)
	default:
		break
	}
}

func convertToOne(filename string) {
	cmd, err := exec.Command(BONNMOTION, "TheONEFile", "-f", filename).Output()
	fmt.Println(string(cmd))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func convertToNSFile(filename string) {
	cmd, err := exec.Command(BONNMOTION, "NSFile", "-f", filename).Output()
	fmt.Println(string(cmd))
	if err != nil {
		fmt.Println(err)
		return
	}
}
func convertToIF(filename string) {
	cmd, err := exec.Command(BONNMOTION, "IntervalFormat", "-f", filename).Output()
	fmt.Println(string(cmd))
	if err != nil {
		fmt.Println(err)
		return
	}
}
