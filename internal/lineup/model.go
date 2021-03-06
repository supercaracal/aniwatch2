package lineup

import (
	"fmt"
	"time"

	"github.com/supercaracal/aniwatch/internal/data"
)

const (
	baseDate = "2020-01-01"
)

func makeLineupsPerDaySlot(dat *data.Data) (*LineupsPerDaySlot, error) {
	v := LineupsPerDaySlot{}

	for _, e := range dat.DayOfWeek {
		v[e.ID] = map[string][]Row{}
		for _, x := range dat.SlotOfDay {
			v[e.ID][x.ID] = []Row{}
		}
	}

	for _, e := range dat.Lineups {
		t, err := time.Parse(time.RFC3339, fmt.Sprintf("%sT%s:00Z", baseDate, e.StartTime))
		if err != nil {
			return nil, err
		}

		l := Row{
			Title:   e.Title,
			Start:   e.StartTime,
			Channel: dat.TvChannel[e.ChannelCode],
		}

		s := dat.GetSlot(t)
		v[e.DayOfWeekCode][s] = append(v[e.DayOfWeekCode][s], l)
	}

	return &v, nil
}

func newIndexData(dat *data.Data, lineups *LineupsPerDaySlot, now time.Time) *IndexData {
	return &IndexData{
		Data:              dat,
		Title:             camelize(dat.AppName),
		Quarter:           fmt.Sprintf("%d-%dQ", now.Year(), calcQuarter(now)),
		LineupCount:       len(dat.Lineups),
		LineupsPerDaySlot: lineups,
	}
}

func camelize(v string) string {
	if v == "" {
		return ""
	}

	if v[0] < 'a' || 'z' < v[0] {
		return v
	}

	return fmt.Sprintf("%s%s", string(v[0]-32), v[1:])
}

func calcQuarter(t time.Time) int {
	var offset int

	if t.Month()%3 == 0 {
		offset = 0
	} else {
		offset = 1
	}

	return int(t.Month())/3 + offset
}
