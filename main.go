package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	_ "embed"

	g "github.com/AllenDang/giu"
)

//go:embed bonnmotion_settings.txt
var theoneConfig string

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

func loop() {

	g.SingleWindow().Layout(
		g.Label("DarMotion"),
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
	wnd := g.NewMasterWindow("DarMotion", 800, 600, 0)
	wnd.Run(loop)
}
