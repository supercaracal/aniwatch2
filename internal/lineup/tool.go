package lineup

import (
	"bytes"
	"time"

	"github.com/supercaracal/aniwatch/internal/data"
)

// GetIndexHTML is
func GetIndexHTML(dat *data.Data) (*bytes.Buffer, error) {
	var buf bytes.Buffer

	tmpl, err := newTemplate()
	if err != nil {
		return nil, err
	}

	lineups, err := makeLineupsPerDaySlot(dat)
	if err != nil {
		return nil, err
	}

	indexData := newIndexData(dat, lineups, time.Now())

	if err := tmpl.render(&buf, "index", indexData); err != nil {
		return nil, err
	}

	return &buf, nil
}
