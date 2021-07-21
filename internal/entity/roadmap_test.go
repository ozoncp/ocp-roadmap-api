package entity

import (
	"testing"
)

func TestRoadmapToString(t *testing.T) {
	roadMap := NewRoadMap(11, 999, "https://some-link.com")
	expected := "Id: 11, UserId: 999, Link: https://some-link.com"
	if expected != roadMap.String() {
		t.Errorf("struct as string must be %q, got %q", expected, roadMap)
	}
}
