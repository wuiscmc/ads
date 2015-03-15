package controllers

import (
	"github.com/wuiscmc/ads/testhelper"
	"testing"
)

func setup() *AdController {
	testhelper.ResetDB()
	return NewAdController()
}

func TestFindAdReturnsMostPrioritizedAd(t *testing.T) {
	controller := setup()
	testhelper.CreateAd("1", "d", 10, 1)
	testhelper.CreateAd("expected", "d", 11, 1)
	ad := controller.FindAd("1")
	if ad.Title != "expected" {
		t.Error("ERROR")
	}
}

func TestFindAdRecordsImpression(t *testing.T) {
	controller := setup()
	testhelper.CreateAd("a", "d", 10, 1)
	controller.FindAd("1")
	controller.FindAd("1")
	impressions := testhelper.ListImpressions()
	if len(impressions) != 2 {
		t.Error("ERROR")
	}
}
