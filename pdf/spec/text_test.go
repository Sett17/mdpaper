package spec

import (
	"testing"
)

func Test_findCutoffSegment(t *testing.T) {
	type args struct {
		segments   []*Segment
		cutoffText string
	}
	tests := []struct {
		name     string
		args     args
		segIdx   int
		location int
	}{
		{"this failed with word method", args{
			segments: []*Segment{
				{Content: "This is the first segment"},
				{Content: "This is the second segment"},
				{Content: ", This is the third segment."},
			},
			cutoffText: "second segment,",
		}, 1, 12},
		{"false positive", args{
			segments: []*Segment{
				{Content: "This is the first segment."},
				{Content: "This is the second segment."},
				{Content: "This is the third segment."},
			},
			cutoffText: "segment. This is the third",
		}, 1, 19},
		{"over 3 segments", args{
			segments: []*Segment{
				{Content: "This is the first segment."},
				{Content: "This is the second segment. It has some text in common with the "},
				{Content: "cutoffText. "},
				{Content: "This is the fourth segment."},
			},
			cutoffText: "in common with the cutoffText. This is",
		}, 1, 45},
		{"Not found", args{
			segments: []*Segment{
				{Content: "This is the first segment."},
				{Content: "This is the second segment."},
				{Content: "This is the third segment."},
			},
			cutoffText: "moin",
		}, -1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := findCutoffSegment(tt.args.segments, tt.args.cutoffText)
			if got != tt.segIdx {
				t.Errorf("findCutoffSegment() segIdx got = %v, want %v", got, tt.segIdx)
			}
			if got1 != tt.location {
				t.Errorf("findCutoffSegment() location  got = %v, want %v", got1, tt.location)
			}
		})
	}
}

func TestEntropy(t *testing.T) {
	//realStrings := []string{
	//	"definite need to test it with a full",
	//	"My specific task was to containerize the",
	//	"hout the phase, my workflow was predom",
	//	"changed a lot during the duration",
	//	" at its core, a way to develop something. I",
	//}

	realStrings := []string{
		"quas",
		"causae",
		"massa",
		"netus",
		"honestatis",
		"mus",
	}

	jumbledStrings := []string{
		"oq=dikka&aqs=chrome..69i57.913j0j1&sourceid=chrome&ie=UTF-8",
		"discrete#EntropyChaoShenBaseE",
		"B086VWXJCM/ref=atv_dp_season_select_s3",
		"T3_2000%20Paper%20Raik%20Rohde.pdf",
	}

	avgRealEnt := 0.0
	biggestRealEnt := 0.0
	for _, s := range realStrings {
		avgRealEnt += Entropy(s)
		if Entropy(s) > biggestRealEnt {
			biggestRealEnt = Entropy(s)
		}
	}
	avgRealEnt /= float64(len(realStrings))

	avgJumbledEnt := 0.0
	smallJumbledEnt := 100.0
	for _, s := range jumbledStrings {
		avgJumbledEnt += Entropy(s)
		if Entropy(s) < smallJumbledEnt {
			smallJumbledEnt = Entropy(s)
		}
	}
	avgJumbledEnt /= float64(len(jumbledStrings))

	t.Logf("Average entropy of real strings: %f", avgRealEnt)
	t.Logf("Biggest entropy of real strings: %f", biggestRealEnt)
	t.Logf("Average entropy of jumbled strings: %f", avgJumbledEnt)
	t.Logf("Smallest entropy of jumbled strings: %f", smallJumbledEnt)

}
