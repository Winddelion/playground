package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int16
	}{{"five words", "Hello I'm definately a human", 5},
		{"three words", "Let it ride", 3},
		{"blank", "", 0},
		{"symbols", "$ % ^ ,,, . '", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWords(tt.input)
			diff := got != tt.want
			if diff {
				t.Errorf("CountWords(%v) = (%v); want %v", tt.input, got, tt.want)
			}
		})
	}

}
