package controllers

import (
	"github.com/wuiscmc/ads/models"
	"github.com/wuiscmc/ads/repositories"
	"strconv"
)

type AdController struct {
	*repositories.AdRepository
	*repositories.EventRepository
}

func NewAdController() *AdController {
	return &AdController{&repositories.AdRepository{}, &repositories.EventRepository{}}
}

func (ac AdController) FindAd(zoneId string) models.Ad {
	zone, _ := strconv.Atoi(zoneId)
	ad := ac.AdRepository.FetchAd(zone)
	ac.AdRepository.LogImpression(ad.Id)
	return ad
}

func (ac AdController) TrackEvent(adId string, action string) bool {
	err := ac.EventRepository.LogEvent(adId, action)
	return err
}

func (ac AdController) SetAdRepository(adRepository *repositories.AdRepository) {
	ac.AdRepository = adRepository
}
