package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/supercaracal/aniwatch/internal/config"
	"github.com/supercaracal/aniwatch/internal/data"
	"github.com/supercaracal/aniwatch/internal/server"
)

const (
	timeout      = 3 * time.Second
	text         = "</body>"
	rootDir      = ".."
	contentDir   = "../docs"
	dataFilePath = "../config/data.yaml"
)

func TestRootPage(t *testing.T) {
	dat, err := data.Load(dataFilePath)
	if err != nil {
		t.Fatal(err)
	}

	mux, err := server.MakeServeMux(config.NewFakeLogger(), dat, rootDir, contentDir)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(mux)
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Error(err)
		return
	}

	hc := &http.Client{Timeout: timeout}
	resp, err := hc.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want=%d, got=%d", http.StatusOK, resp.StatusCode)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}

	if body := string(bytes); !strings.Contains(body, text) {
		t.Errorf("'%s' is not found in\n```\n%s```\n", text, body)
	}
}