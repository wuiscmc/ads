package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wuiscmc/ads/controllers"
	"github.com/wuiscmc/ads/models"
	"github.com/wuiscmc/ads/repositories"
	"log"
	"net/http"
	"os"
	"text/template"
)

var Logger = log.New(os.Stdout, "[SERVER]", 0)

type Response struct {
	models.Ad
	Host string
}

func main() {

	router := httprouter.New()
	adRepository := repositories.NewAdRepository()
	adController := controllers.NewAdController(adRepository)

	router.GET(v1("/ads/:zoneId"), func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logRequest(r)
		ad := adController.FindAd(ps.ByName("zoneId"))
		tmpl, _ := template.ParseFiles("templates/advast2.tmpl")
		w.Header().Set("Content-Type", "application/xml")
		tmpl.Execute(w, Response{ad, string("http://localhost:3000")})
	})

	router.POST("/track", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logRequest(r)
		query := r.URL.Query()
		id := query.Get("id")
		action := query.Get("action")
		adController.TrackEvent(id, action)
	})
	http.ListenAndServe(":3000", router)
}

func v1(route string) string {
	return string("/v1" + route)
}

func logRequest(r *http.Request) {
	Logger.Printf(" %s %s", r.Method, r.URL.Path)
}
