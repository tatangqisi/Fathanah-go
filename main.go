package main

import (
	"kammi/controllers/ahd"
	"kammi/controllers/articlecontroller"
	"kammi/controllers/asetcontroller"
	"kammi/controllers/authcontroller"
	"kammi/controllers/dhariancontroller"
	"kammi/controllers/diarycontroller"
	"kammi/controllers/homecontroller"
	"kammi/controllers/kidscontroller"
	"kammi/controllers/mssgcontroller"
	"kammi/controllers/productcontroller"
	"kammi/controllers/qurancontroller"
	"kammi/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/Img", authcontroller.Imgp).Methods("GET")
	r.HandleFunc("/home", homecontroller.Home).Methods("GET")
	// r.HandleFunc("/home/{id}", homecontroller.Homelog).Methods("GET")
	r.HandleFunc("/logo-icon", homecontroller.Header).Methods("GET")

	r.HandleFunc("/ashusna", ahd.Ashusna).Methods("GET")
	r.HandleFunc("/dharian", dhariancontroller.KDharian).Methods("GET")
	r.HandleFunc("/dharian/{menu}", dhariancontroller.Show).Methods("GET")
	r.HandleFunc("/dharian/{menu}", dhariancontroller.Show).Methods("POST")
	r.HandleFunc("/articlee", articlecontroller.Kberita).Methods("GET")
	r.HandleFunc("/article", articlecontroller.Showm).Methods("GET")
	r.HandleFunc("/article", articlecontroller.Showm).Methods("POST")
	r.HandleFunc("/article/{id}", articlecontroller.Show).Methods("GET")
	r.HandleFunc("/kids", kidscontroller.Mkids).Methods("GET")
	r.HandleFunc("/kids/{menu}", kidscontroller.Show).Methods("GET")
	r.HandleFunc("/quran", qurancontroller.Qrnsurah).Methods("GET")
	r.HandleFunc("/quran", qurancontroller.Qrnsurah).Methods("POST")
	r.HandleFunc("/quran/{surah}", qurancontroller.Shows).Methods("GET")

	r.HandleFunc("/products", productcontroller.Index).Methods("GET")

	r.HandleFunc("/sign/{id}", authcontroller.Show).Methods("GET")
	r.HandleFunc("/sign/{id}/update", authcontroller.Updateprofile).Methods("POST")
	r.HandleFunc("/sign/{id}/updateuser", authcontroller.Updateusername).Methods("POST")
	r.HandleFunc("/sign/{id}/sun", authcontroller.Showun).Methods("GET")
	r.HandleFunc("/sign/{id}/updateun", authcontroller.Updateusername).Methods("POST")
	r.HandleFunc("/sign/{id}/updatepw", authcontroller.Updatepw).Methods("POST")

	r.HandleFunc("/sign/{user}/diary", diarycontroller.GetAllDiary).Methods("GET")
	r.HandleFunc("/sign/{user}/diary/create", diarycontroller.CreateDiary).Methods("POST")
	r.HandleFunc("/sign/{user}/diary/{no}", diarycontroller.GetDiary).Methods("GET")
	r.HandleFunc("/sign/{user}/diary/{no}/update", diarycontroller.UpdateDiary).Methods("POST")
	r.HandleFunc("/sign/{user}/diary/{no}/delete", diarycontroller.DeleteDiary).Methods("POST")

	r.HandleFunc("/sign/{id}/message", mssgcontroller.Message).Methods("POST")

	r.HandleFunc("/about", asetcontroller.About).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", r))
}
