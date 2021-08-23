package entity

import (
	"testing"
	"time"
)

func TestRoadmapToString(t *testing.T) {
	roadMap := NewRoadMap(11, 999, "https://some-link.com", time.Now())
	expected := "Id: 11, UserId: 999, Link: https://some-link.com"
	if expected != roadMap.String() {
		t.Errorf("struct as string must be %q, got %q", expected, roadMap)
	}
}
