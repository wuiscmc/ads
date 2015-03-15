package controllers

import (
	"github.com/wuiscmc/ads/models"
	"github.com/wuiscmc/ads/repositories"
)

type AdController struct {
	*repositories.AdRepository
	*repositories.EventRepository
}

func NewAdController(ar *repositories.AdRepository) *AdController {
	return &AdController{ar, &repositories.EventRepository{}}
}

func (ac AdController) FindAd(zoneId string) models.Ad {
	ad := ac.AdRepository.FetchAd(zoneId)
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
