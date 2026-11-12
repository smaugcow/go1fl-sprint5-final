package spentenergy

import (
	"testing"
	"time"
)

func TestMeanSpeedNegativeSteps(t *testing.T) {
	if speed := MeanSpeed(-10, 1.75, time.Hour); speed != 0 {
		t.Fatalf("negative steps: speed=%v", speed)
	}
}
