package main

import "testing"

func TestCelciusToFarenheit(t *testing.T) {
	const epsilon = 1e-5
	tests := []struct {
		name  string
		input float32
		want  float32
	}{{"freezing", 0, 32},
		{"boiling", 100, 212},
		{"body", 37, 98.6},
		{"fractional", 36.6, 97.88},
		{"negative", -40, -40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelciusToFarenheit(tt.input)
			diff := got - tt.want
			if diff < 0 {
				diff = -diff
			}
			if diff > epsilon {
				t.Errorf("CelciusToFarenheit(%v) = (%v); want %v (+-%v)",
					tt.input, got, tt.want, epsilon)
			}
		})
	}
}
