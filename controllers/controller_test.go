package controllers

import (
	"github.com/wuiscmc/ads/testhelper"
	"testing"
)

func newAdController() *AdController {
	testhelper.ResetDB()
	return NewAdController()
}

func TestFindAdReturnsMostPrioritizedAd(t *testing.T) {
	controller := newAdController()
	testhelper.CreateAd("1", "d", 10, 1)
	testhelper.CreateAd("expected", "d", 11, 1)
	ad := controller.FindAd("1")
	if ad.Title != "expected" {
		t.Error("Returned Ad doesnt isnt the expected one")
	}
}

func TestFindAdReturnsNoAd(t *testing.T) {
	controller := newAdController()
	ad := controller.FindAd("2")
	if ad.Id == 0 && ad.Title != "" {
		t.Error("Returned Ad doesnt isnt the expected one")
	}
}

func TestFindAdRecordsImpression(t *testing.T) {
	controller := newAdController()
	testhelper.CreateAd("a", "d", 10, 1)
	controller.FindAd("1")
	controller.FindAd("1")
	impressions := testhelper.ListImpressions()
	if len(impressions) != 2 {
		t.Error("Expected 2 impresssions in the DB")
	}
}

func TestTrackEvent(t *testing.T) {
	controller := newAdController()
	testhelper.CreateAd("a", "d", 10, 1)
	controller.TrackEvent("1", "vivaelabuelo")
	event := testhelper.FetchLastEventTracked()
	if event.Action != "vivaelabuelo" {
		t.Error("The tracked event doesnt correspond to the Ad")
	}
}