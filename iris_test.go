package iris

import "testing"

func TestRgb(t *testing.T) {
	tests := []struct {
		name  string
		input string
		wantR int
		wantG int
		wantB int
	}{
		{
			name:  "valid hex with hash",
			input: "#FF0000",
			wantR: 255,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "valid hex without hash",
			input: "00FF00",
			wantR: 0,
			wantG: 255,
			wantB: 0,
		},
		{
			name:  "black",
			input: "000000",
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "white",
			input: "FFFFFF",
			wantR: 255,
			wantG: 255,
			wantB: 255,
		},
		{
			name:  "invalid length",
			input: "FF00",
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
		{
			name:  "invalid characters",
			input: "GGHHII",
			wantR: 0,
			wantG: 0,
			wantB: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := rgb(tt.input)
			if gotR != tt.wantR || gotG != tt.wantG || gotB != tt.wantB {
				t.Errorf("rgb(%q) = (%v, %v, %v), want (%v, %v, %v)",
					tt.input, gotR, gotG, gotB, tt.wantR, tt.wantG, tt.wantB)
			}
		})
	}

	text := SetColor("hello world", "#FF0000", "#00000")
	stripped := Strip(text)
	if stripped != "hello world" {
		t.Errorf("expected 'hello world' but got %s", stripped)
	}
}
