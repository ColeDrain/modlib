package modlib

import "testing"

func TestConfig(t *testing.T) {
	var tests []struct {
		name string
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Config(); got != tt.want {
				t.Errorf("Config() = %v, want %v", got, tt.want)
			}
		})
	}
}
