package homecontroller

import (
	"encoding/json"
	"kammi/models"
	"log"
	"net/http"

)

func Home(w http.ResponseWriter, r *http.Request) {
	sld := []models.Slide{}
	Ftr := []models.Feature{}
	brt := []models.Vbrt{}
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var response models.Home
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if logo != nil {
		log.Print(logo.Error())
	}
	result1 := models.DB.Table("web-slider").Select("`web-slider`.`id`, `web-slider`.`name`, `web-slider`.`img`, `img-path`.`path`, `web-slider`.`status`, `web-slider`.`url`").Joins("INNER JOIN `img-path` ON `web-slider`.`path` =`img-path`.`id`").Find(&sld).Error
	if result1 != nil {
		log.Print(result1.Error())
	}
	result2 := models.DB.Table("web-feature").Select("`web-feature`.`id`, `web-feature`.`name`, `web-feature`.`img`, `img-path`.`path`, `web-feature`.`url`").Joins("INNER JOIN `img-path` ON `web-feature`.`path` =`img-path`.`id`").Find(&Ftr).Error
	if result2 != nil {
		log.Print(result2.Error())
	}
	result3 := models.DB.Table("article-data").Select("`article-data`.`id`, `article-data`.`time`, `article-data`.`img`, `article-data`.`title`, `article-category`.`category`, `article-data`.`desc`, `img-path`.`path`").Joins("INNER JOIN `img-path` JOIN `article-category` ON `article-data`.`path` =`img-path`.`id` AND `article-data`.`category` =`article-category`.`id`").Limit(5).Order("time DESC").Find(&brt).Error
	if result3 != nil {
		log.Print(result3.Error())
	}
	response.Icon = Ic
	response.Logo = Lg
	response.Data1 = sld
	response.Data2 = Ftr
	response.Data3 = brt
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

//logo dan icon
func Header(w http.ResponseWriter, _ *http.Request) {
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var response models.Hd
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if logo != nil {
		log.Print(logo.Error())
	}
	response.Icon = Ic
	response.Logo = Lg
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
