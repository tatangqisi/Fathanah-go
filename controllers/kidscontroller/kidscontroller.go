package kidscontroller

import (
	"encoding/json"
	"kammi/helper"
	"kammi/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

)

func Mkids(w http.ResponseWriter, r *http.Request) {
	Jk := []models.Jkids{}
	var response models.Jkd
	result := models.DB.Table("kidz-menu").Scan(&Jk).Error
	if result != nil {
		log.Print(result.Error())
	}
	response.Data = Jk
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

//show berita controller
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	menu, err := strconv.ParseInt(vars["menu"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var dh []models.Ikids
	var Ddh models.Dkds
	// dhInput := models.Dh{Katagori: int(katagori)}
	if err := models.DB.Table("kidz-data").Select("`kidz-data`.`title`, `kidz-data`.`desc`, `kidz-data`.`img`, `kidz-menu`.`menu`").Joins("INNER JOIN `kidz-menu` ON `kidz-data`.`menu` = `kidz-menu`.`id`").Where("`kidz-data`.`menu` = ?", menu).Find(&dh).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	Ddh.Data = dh
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(Ddh)
}
