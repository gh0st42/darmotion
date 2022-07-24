package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	_ "embed"

	g "github.com/AllenDang/giu"
)

//go:embed bonnmotion_settings.txt
var theoneConfig string

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

func onRun() {
	fmt.Println("Run")

	var fout string
	if filename == "" {
		fout = "move"
	} else {
		fout = filename
	}
	convertToOne(fout)
	one_cfg := strings.Replace(theoneConfig, "Scenario.endTime = DURATION", fmt.Sprintf("Scenario.endTime = %v", duration), 1)
	one_cfg = strings.Replace(one_cfg, "NODES", fmt.Sprintf("%v", num_nodes), 1)
	one_cfg = strings.Replace(one_cfg, "WIDTH", fmt.Sprintf("%v", width), 1)
	one_cfg = strings.Replace(one_cfg, "HEIGHT", fmt.Sprintf("%v", height), 1)
	one_cfg = strings.Replace(one_cfg, "FILENAME", fmt.Sprintf("%v.one", fout), 1)

	//fmt.Println(one_cfg)
	file, err := os.Create("one_settings.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.WriteString(one_cfg)
	file.Close()

	fmt.Println(file.Name())

	cmd := exec.Command("one.sh", file.Name())
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
}

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
)

func loop() {

	g.SingleWindow().Layout(
		g.Label("TuDarMotion"),
		g.InputText(&filename).Label("filename"),
		g.Combo("format", formats[selected], formats, &selected),
		g.Spacing(),
		g.Row(
			g.Checkbox("batch", &batch),
			g.InputInt(&reps).Size(100).Label("reptitions"),
		),
		g.Spacing(),
		g.Separator(),
		g.Spacing(),
		g.InputInt(&num_nodes).Size(100).Label("nodes"),
		g.InputInt(&duration).Size(100).Label("duration in seconds"),
		g.InputInt(&skip).Size(100).Label("seconds to skip"),
		g.Row(
			g.InputInt(&width).Size(100),
			g.Label("x"),
			g.InputInt(&height).Size(100).Label("area in m^2"),
		),
		g.Spacing(),
		g.TabBar().TabItems(
			g.TabItem("Random Waypoint").Layout(
				g.Spacing(),
				g.Label("Random Waypoint Model"),
				g.Row(
					g.Button("Generate").OnClick(onGenerateRandomWaypoint),
					g.Button("Run").OnClick(onRun),
				),
				g.InputTextMultiline(&multiline).Size(g.Auto, g.Auto),
			),
			g.TabItem("SMOOTH").Layout(
				g.Spacing(),
				g.Label("SMOOTH Mobility Model"),
				g.Row(
					g.Button("Generate").OnClick(onGenerateSMOOTH),
					g.Button("Run").OnClick(onRun),
				),
				g.InputInt(&smooth.comm_range).Size(100).Label("range"),
				g.InputInt(&smooth.clusters).Size(100).Label("clusters"),
				g.InputFloat(&smooth.alpha).Size(100).Label("alpha"),
				g.InputInt(&smooth.f_min).Size(100).Label("f_min"),
				g.InputInt(&smooth.f_max).Size(100).Label("f_max"),
				g.InputFloat(&smooth.beta).Size(100).Label("beta"),
				g.InputInt(&smooth.p_min).Size(100).Label("p_min"),
				g.InputInt(&smooth.p_max).Size(100).Label("p_max"),
			),
		),
	)
}

func main() {
	wnd := g.NewMasterWindow("TuDarMotion", 800, 600, 0)
	wnd.Run(loop)
}
