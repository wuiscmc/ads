package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/wuiscmc/ads/controllers"
	"github.com/wuiscmc/ads/models"
	"log"
	"net/http"
	"os"
	"text/template"
)

var logger = log.New(os.Stdout, "[SERVER]", 0)

var adController = controllers.NewAdController()

type Response struct {
	models.Ad
	Host string
}

func main() {
	router := httprouter.New()
	router.GET(v1("/ads/:zoneId"), fetchAdHandler)
	router.POST(v1("/track"), trackEventHandler)
	http.ListenAndServe(":3000", router)
}

func fetchAdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
	ad := adController.FindAd(ps.ByName("zoneId"))
	if ad.Id == 0 {
		w.WriteHeader(404)
	} else {
		tmpl, _ := template.ParseFiles("templates/advast2.tmpl")
		w.Header().Set("Content-Type", "application/xml")
		tmpl.Execute(w, Response{ad, string("http://localhost:3000/v1")})
	}
}

func trackEventHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
	query := r.URL.Query()
	id := query.Get("uid")
	action := query.Get("action")
	adController.TrackEvent(id, action)
	w.WriteHeader(204)
}

func v1(route string) string {
	return string("/v1" + route)
}

func logRequest(r *http.Request) {
	logger.Printf(" %s %s", r.Method, r.URL.Path)
}
