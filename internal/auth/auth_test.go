package auth

import "testing"

func TestGetAuth(t *testing.T) {
	var tests []struct {
		name string
		want string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAuth(); got != tt.want {
				t.Errorf("GetAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
