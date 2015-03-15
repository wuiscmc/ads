package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wuiscmc/ads/testhelper"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFetchAdHandler(t *testing.T) {
	testhelper.SetTestMode()
	testhelper.CreateAd("test", "testFetchAdHandler", 1, 1)

	req, _ := http.NewRequest("GET", "http://whatever.com/ads/1", nil)
	w := httptest.NewRecorder()

	fetchAdHandler(w, req, httprouter.Params{httprouter.Param{"zoneId", "1"}})

	if w.Header().Get("Content-Type") != "application/xml" {
		t.Error("Doesnt return XML")
	}

	if !strings.Contains(w.Body.String(), "VAST") {
		t.Error("Invalid Format")
	}
}

func TestFetchAdHandlerUnexistingZone(t *testing.T) {
	testhelper.SetTestMode()

	req, _ := http.NewRequest("GET", "http://whatever.com/ads/1", nil)
	w := httptest.NewRecorder()
	fetchAdHandler(w, req, httprouter.Params{httprouter.Param{"zoneId", "1231"}})
	if w.Code != 404 {
		t.Error("The Zone doesnt exist therefore it's a not found resource")
	}
}

func TestTrackEventHandler(t *testing.T) {
	testhelper.SetTestMode()

	req, _ := http.NewRequest("POST", "http://whatever.com/track", nil)
	w := httptest.NewRecorder()

	trackEventHandler(w, req, httprouter.Params{})

	if w.Code != 204 {
		t.Error("Return HTTP code isnt 204")
	}
}
