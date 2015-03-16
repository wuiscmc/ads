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
	http.ListenAndServe(getPort(), router)
}

func fetchAdHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
	ad := adController.FindAd(ps.ByName("zoneId"))
	setCorsHeaders(w)
	if ad.Id == 0 {
		w.WriteHeader(404)
	} else {
		tmpl, _ := template.ParseFiles("templates/advast2.tmpl")
		w.Header().Set("Content-Type", "application/xml")
		tmpl.Execute(w, Response{ad, string(getHost() + "/v1")})
	}
}

func trackEventHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
	query := r.URL.Query()
	id := query.Get("uid")
	action := query.Get("action")
	adController.TrackEvent(id, action)
	setCorsHeaders(w)
	w.WriteHeader(204)
}

func setCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func v1(route string) string {
	return string("/v1" + route)
}

func logRequest(r *http.Request) {
	logger.Printf(" %s %s", r.Method, r.URL.Path)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
		logger.Printf("Defaulting to port " + port)
	}
	return ":" + port
}

func getHost() string {
	if os.Getenv("ENV") == "production" {
		return os.Getenv("HOSTNAME")
	} else {
		return "http://localhost:3000"
	}
}
