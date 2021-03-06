package server

import (
	"fmt"
	"net/http"

	"github.com/supercaracal/aniwatch/internal/config"
	"github.com/supercaracal/aniwatch/internal/data"
	"github.com/supercaracal/aniwatch/internal/lineup"
	"github.com/supercaracal/aniwatch/internal/middleware"
)

// MakeServeMux is
func MakeServeMux(logger *config.Logger, dat *data.Data, contentDir string) (http.Handler, error) {
	mux := http.NewServeMux()

	if err := setUpRoot(mux, logger, dat); err != nil {
		return nil, err
	}

	setUpStaticContents(mux, dat, contentDir)

	return setUpMiddlewares(mux, logger), nil
}

func setUpMiddlewares(h http.Handler, logger *config.Logger) http.Handler {
	h = middleware.AccessLog(h, logger.Info)
	h = middleware.CompressResponse(h)
	return h
}

func setUpRoot(mux *http.ServeMux, logger *config.Logger, dat *data.Data) error {
	ctrl, err := lineup.NewController(dat, logger)
	if err != nil {
		return err
	}

	mux.HandleFunc("/", ctrl.Exec)

	return nil
}

func setUpStaticContents(mux *http.ServeMux, dat *data.Data, contentDir string) {
	files := []string{"favicon.ico", "style.css"}

	for _, f := range files {
		mux.HandleFunc(
			fmt.Sprintf("/%s/%s", dat.AppName, f),
			func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, fmt.Sprintf("%s/%s", contentDir, f))
			},
		)
	}
}
