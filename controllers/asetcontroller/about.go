package asetcontroller

import (
	"encoding/json"
	"fathanah/models"
	"log"
	"net/http"

)

func About(w http.ResponseWriter, _ *http.Request) {
	ab := []models.Ab{}
	ab2 := []models.Ab{}
	team := []models.Team{}
	Lg := []models.Head{}
	var response models.Dab
	Data := models.DB.Table("about-data").Select("`about-data`.`id`, `about-data`.`desc`, `about-data`.`name`, `about-data`.`img`, `img-asset`.`path`").Joins("INNER JOIN `img-asset` ON `about-data`.`img` =`img-asset`.`id`").Where("`about-data`.`id` = 1").Scan(&ab).Error
	if Data != nil {
		log.Print(Data.Error())
	}
	Data2 := models.DB.Table("about-data").Select("`about-data`.`id`,`about-data`.`desc`, `about-data`.`name`, `about-data`.`img`, `img-asset`.`path`").Joins("INNER JOIN `img-asset` ON `about-data`.`img` =`img-asset`.`id`").Where("`about-data`.`id` = 2").Scan(&ab2).Error
	if Data2 != nil {
		log.Print(Data.Error())
	}
	Data3 := models.DB.Table("admin-data").Scan(&team).Error
	if Data3 != nil {
		log.Print(Data.Error())
	}
	Logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if Logo != nil {
		log.Print(Logo.Error())
	}
	//icon
	Ic := []models.Icon{}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	response.Icon = Ic
	response.Logo = Lg
	response.Data1 = ab
	response.Data2 = ab2
	response.Data3 = team
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
