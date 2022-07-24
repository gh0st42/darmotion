package main

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
