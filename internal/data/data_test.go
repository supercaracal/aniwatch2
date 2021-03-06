package data

import (
	"testing"
	"time"
)

func TestGetSlot(t *testing.T) {
	cases := []struct {
		timeStr string
		want    string
	}{
		{"2020-01-01T00:00:00Z", "midnight"},
		{"2020-01-01T04:59:59Z", "midnight"},
		{"2020-01-01T05:00:00Z", "daytime"},
		{"2020-01-01T15:59:59Z", "daytime"},
		{"2020-01-01T16:00:00Z", "sunset"},
		{"2020-01-01T18:59:59Z", "sunset"},
		{"2020-01-01T19:00:00Z", "night"},
		{"2020-01-01T23:59:59Z", "night"},
	}

	dt, err := Load()
	if err != nil {
		t.Fatal(err)
	}

	for n, c := range cases {
		argTime, err := time.Parse(time.RFC3339, c.timeStr)
		if err != nil {
			t.Fatal(err)
		}

		if got := dt.GetSlot(argTime); c.want != got {
			t.Errorf("%d: want=%s, got=%s", n+1, c.want, got)
		}
	}
}
