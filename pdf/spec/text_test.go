package spec

import "testing"

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
		{"over single segment", args{
			segments: []*Segment{
				{Content: "This is the first segment."},
				{Content: "This is the second segment."},
				{Content: "This is the third segment."},
			},
			cutoffText: "second segment.",
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
			cutoffText: "second segment. It has some text in common with the cutoffText. This is",
		}, 1, 12},
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
