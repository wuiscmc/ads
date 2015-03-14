package controllers

import (
	"github.com/wuiscmc/ads/models"
	"github.com/wuiscmc/ads/repositories"
)

type AdController struct {
	adRepository *repositories.AdRepository
}

func NewAdController(adRepository *repositories.AdRepository) *AdController {
	return &AdController{adRepository}
}

func (ac AdController) FindAd(zoneId string) models.Ad {
	ad := ac.adRepository.FetchAd(zoneId)
	ac.adRepository.LogImpression(ad.Id)
	return ad
}

func (ac AdController) Track() {
}
