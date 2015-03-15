package controllers

import (
	"github.com/wuiscmc/ads/testhelper"
	"testing"
)

func Setup() *AdController {
	testhelper.ResetDB()
	return NewAdController(testhelper.NewAdRepository())
}

func TestFindAdReturnsMostPrioritizedAd(t *testing.T) {
	controller := Setup()
	testhelper.CreateAd("1", "a", "d", 10, 1)
	testhelper.CreateAd("2", "b", "d", 11, 1)
	ad := controller.FindAd("1")
	if ad.Id != "2" {
		t.Error("ERROR")
	}
}

func TestFindAdRecordsImpression(t *testing.T) {
	controller := Setup()
	testhelper.CreateAd("1", "a", "d", 10, 1)
	controller.FindAd("1")
	controller.FindAd("1")
	impressions := testhelper.ListImpressions()
	if len(impressions) != 2 {
		t.Error("ERROR")
	}
}
