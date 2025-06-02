package cueparser

import (
	"github.com/google/go-cmp/cmp"
	"os"
	"path/filepath"
	"testing"
)

func TestParseCueFiles(t *testing.T) {
	tests := []struct {
		name             string
		filename         string
		expectedCueSheet CueSheet
	}{
		{
			name:     "Test #1 - Dark Side of the Moon",
			filename: "The Dark Side of the Moon.cue",
			expectedCueSheet: CueSheet{
				Title:     "The Dark Side of the Moon",
				Performer: "Pink Floyd",
				Files: []File{
					{
						Name: "01 - Speak to Me.flac",
						Type: "WAVE",
						Tracks: []Track{
							{
								Number:    1,
								Title:     "Speak to Me",
								Performer: "Pink Floyd",
								Indexes: []Index{
									{
										Number: 1,
										Time:   "00:00:00",
									},
								},
							},
							{
								Number:    2,
								Title:     "Breathe",
								Performer: "Pink Floyd",
								Indexes: []Index{
									{
										Number: 0,
										Time:   "01:02:34",
									},
									{
										Number: 1,
										Time:   "01:08:00",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.Open(filepath.Join("testdata", tt.filename))
			if err != nil {
				t.Fatalf("failed to open test cue file: %v", err)
			}
			defer f.Close()

			sheet, err := Parse(f)
			if err != nil {
				t.Fatalf("failed to parse cue file: %v", err)
			}

			if sheet == nil {
				t.Fatalf("nil parsing result returned, withot any errors")
			}

			if diff := cmp.Diff(tt.expectedCueSheet, *sheet); diff != "" {
				t.Errorf("CueSheet mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
