package conf

import (
	"os"
	"testing"
)

func TestWorkDir(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "get work dir by env MDCLUBGO_WORK_DIR",
			env:  "/tmp",
			want: "/tmp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("MDCLUBGO_WORK_DIR", tt.env)

			if got := WorkDir(); got != tt.want {
				t.Errorf("WorkDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomDir(t *testing.T) {
	tests := []struct {
		name string
		env  string
		want string
	}{
		{
			name: "should append work dir and '/custom'",
			env:  "/tmp",
			want: "/tmp/custom",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Setenv("MDCLUBGO_WORK_DIR", tt.env)

			if got := CustomDir(); got != tt.want {
				t.Errorf("CustomDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
