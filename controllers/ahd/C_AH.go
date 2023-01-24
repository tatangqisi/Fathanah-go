package ahd

import (
	"encoding/json"
	"fathanah/models"
	"log"
	"net/http"

)

func Ashusna(w http.ResponseWriter, r *http.Request) {
	husna := []models.AH{}
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var response models.R_ah
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	logo := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if logo != nil {
		log.Print(logo.Error())
	}
	result := models.DB.Table("asmaul-husna-data").Scan(&husna).Error
	if result != nil {
		log.Print(result.Error())
	}
	response.Icon = Ic
	response.Logo = Lg
	response.Data = husna
	// ash, err := json.Marshal(response)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// test := helper.Encrypt(string(ash))
	// response2 := map[string]string{"data": test}
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}
