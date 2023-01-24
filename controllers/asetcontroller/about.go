package asetcontroller

import (
	"encoding/json"
	"kammi/models"
	"log"
	"net/http"

)

func About(w http.ResponseWriter, _ *http.Request) {
	ab := []models.Ab{}
	ab2 := []models.Ab{}
	team := []models.Team{}
	Lg := []models.Head{}
	var response models.Dab
	result := models.DB.Table("about-data").Select("`about-data`.`id`, `about-data`.`desc`, `about-data`.`name`, `about-data`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `about-data`.`path` =`img-path`.`id`").Where("`about-data`.`id` = 1").Scan(&ab).Error
	if result != nil {
		log.Print(result.Error())
	}
	result2 := models.DB.Table("about-data").Select("`about-data`.`id`,`about-data`.`desc`, `about-data`.`name`, `about-data`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `about-data`.`path` =`img-path`.`id`").Where("`about-data`.`id` = 2").Scan(&ab2).Error
	if result2 != nil {
		log.Print(result.Error())
	}
	result3 := models.DB.Table("web-team").Scan(&team).Error
	if result3 != nil {
		log.Print(result.Error())
	}
	result4 := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result4 != nil {
		log.Print(result4.Error())
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
