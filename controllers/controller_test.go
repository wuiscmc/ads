package controllers

import (
	"github.com/wuiscmc/ads/testhelper"
	"testing"
)

func Setup() *AdController {
	testhelper.ResetDB()
	testhelper.SeedDB()
	return NewAdController(testhelper.NewAdRepository())
}

func TestFindAdReturnsMostPrioritizedAd(t *testing.T) {
	controller := Setup()
	ad := controller.FindAd("1")
	if ad.ZoneField != "1" {
		t.Error("ERROR")
	}
}

func TestFindAdRecordsImpression(t *testing.T) {
	controller := Setup()
	impressionsBefore := testhelper.ListImpressions()
	controller.FindAd("1")
	impressionsAfter := testhelper.ListImpressions()
	lengthAfter := len(impressionsAfter)
	if len(impressionsBefore) != lengthAfter-1 {
		t.Error("ERROR")
	}

}
