package articlecontroller

import (
	"encoding/json"
	"fathanah/helper"
	"fathanah/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

)

func Kberita(w http.ResponseWriter, _ *http.Request) {
	kb := []models.Cart{}
	Lg := []models.Head{}
	var response models.Kdb
	result := models.DB.Table("article-category").Scan(&kb).Error
	if result != nil {
		log.Print(result.Error())
	}
	result1 := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result1 != nil {
		log.Print(result1.Error())
	}
	Ic := []models.Icon{}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	response.Icon = Ic
	response.Logo = Lg
	response.Data = kb
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

//show berita controller
func Showm(w http.ResponseWriter, r *http.Request) {
	var brt []models.Art
	Lg := []models.Head{}
	Ac := []models.Cart{}
	Ic := []models.Icon{}
	var Ddh models.Dart
	title := r.FormValue("title")
	category := models.DB.Table("article-category").Scan(&Ac).Error
	if category != nil {
		log.Print(category.Error())
	}
	if err := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`time`, `article-data`.`img`, `article-data`.`title`, `article-category`.`category`, `article-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `article-category` JOIN `img-path` ON `article-data`.`category` = `article-category`.`id` AND `article-data`.`path` =`img-path`.`id`").Where("`article-data`.`title` LIKE ?", fmt.Sprintf("%%%s%%", title)).Find(&brt).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	result := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result != nil {
		log.Print(result.Error())
	}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	Ddh.Icon = Ic
	Ddh.Logo = Lg
	Ddh.Category = Ac
	Ddh.Data = brt
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(Ddh)
}

func Showc(w http.ResponseWriter, r *http.Request) {
	var brt []models.Art
	Lg := []models.Head{}
	Pac := []models.Cart{}
	Ac := []models.Cart{}
	Ic := []models.Icon{}
	var Ddh models.Cartl
	vars := mux.Vars(r)
	title := r.FormValue("title")
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	categorypick := models.DB.Table("article-category").Where("id = ?", id).Scan(&Pac).Error
	if categorypick != nil {
		log.Print(categorypick.Error())
	}
	category := models.DB.Table("article-category").Scan(&Ac).Error
	if category != nil {
		log.Print(category.Error())
	}
	if err := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`time`, `article-data`.`img`, `article-data`.`title`, `article-category`.`category`, `article-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `article-category` JOIN `img-path` ON `article-data`.`category` = `article-category`.`id` AND `article-data`.`path` =`img-path`.`id`").Where("`article-data`.`title` LIKE ?", fmt.Sprintf("%%%s%%", title)).Where("`article-category`.`id` = ?", id).Find(&brt).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	result := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result != nil {
		log.Print(result.Error())
	}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	Ddh.Icon = Ic
	Ddh.Logo = Lg
	Ddh.PickC = Pac
	Ddh.Category = Ac
	Ddh.Data = brt
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(Ddh)
}

//show berita controller
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var brt []models.Art
	var brtn []models.Artn
	Ic := []models.Icon{}
	Lg := []models.Head{}
	var Ddh models.Vart
	if err := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`time`, `article-data`.`img`, `article-data`.`title`, `article-category`.`category`, `article-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `article-category` JOIN `img-path` ON `article-data`.`category` = `article-category`.`id` AND `article-data`.`path` =`img-path`.`id`").Where("`article-data`.`id`= ?", id).Find(&brt).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	result2 := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`title`, `article-data`.`time`").Limit(5).Order("time DESC").Find(&brtn).Error
	if result2 != nil {
		log.Print(result2.Error())
	}
	logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if logo != nil {
		log.Print(logo.Error())
	}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	Ddh.Icon = Ic
	Ddh.Logo = Lg
	Ddh.Data = brt
	Ddh.Data2 = brtn
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(Ddh)
}
