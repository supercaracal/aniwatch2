package lineup

import (
	"fmt"
	"net/http"
	"time"

	"github.com/supercaracal/aniwatch/internal/config"
	"github.com/supercaracal/aniwatch/internal/data"
)

// Controller is
type Controller struct {
	data   *data.Data
	logger *config.Logger
	tmpl   *tmplobj
}

// NewController is
func NewController(dat *data.Data, logger *config.Logger) (*Controller, error) {
	tmpl, err := newTemplate()
	if err != nil {
		return nil, err
	}

	return &Controller{
		data:   dat,
		logger: logger,
		tmpl:   tmpl,
	}, nil
}

// Exec is
func (ctrl *Controller) Exec(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ctrl.index(w, r)
		return
	}

	http.NotFound(w, r)
}

func (ctrl *Controller) index(w http.ResponseWriter, r *http.Request) {
	lineups, err := buildLineupsPerDaySlot(ctrl.data)
	if err != nil {
		ctrl.logger.Err.Println(fmt.Errorf("Failed to build lineups: %w", err))
		responseInternalServerError(w)
		return
	}

	indexData := newIndexData(ctrl.data, lineups, time.Now())
	if err := ctrl.tmpl.render(w, "index", indexData); err != nil {
		ctrl.logger.Err.Println(fmt.Errorf("Failed to render html file: index: %w", err))
		responseInternalServerError(w)
	}
}

func responseInternalServerError(w http.ResponseWriter) {
	msg := fmt.Sprintf("%d internal server error", http.StatusInternalServerError)
	http.Error(w, msg, http.StatusInternalServerError)
}
