package diarycontroller

import (
	"encoding/json"
	"fathanah/helper"
	"fathanah/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

)

func GetAllDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var diary []models.Diaryscn
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var response models.Dd
	diaryInput := models.Diary{User: user}
	if err := models.DB.Table("diary-data").Where(&diaryInput, user).Order("time DESC").Find(&diary-data).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	result := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result != nil {
		log.Print(result.Error())
	}
	response.Data = diary
	response.Logo = Lg
	response.Icon = Ic
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(response)
}

func CreateDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errr := r.ParseForm()
	if errr != nil {
		panic(errr)
	}
	//mengambil user dari parameter
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// memasukan parameter user ke database dan membuat datetime
	now := time.Now()
	diaryInput := models.Diary{User: user, Time: now.String()}
	subject := r.Form.Get("subject")
	body := r.Form.Get("body")
	diaryInput.Subject = subject
	diaryInput.Body = body

	// input ke database
	if err := models.DB.Table("diary-data").Create(&diaryInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	diaryInput.User = user

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func GetDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errr := r.ParseForm()
	if errr != nil {
		panic(errr)
	}
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	no, err := strconv.ParseInt(vars["no"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var diary []models.Diarys
	Lg := []models.Head{}
	Ic := []models.Icon{}
	var response models.Dd2
	diaryInput := models.Diary{User: user, No: no}

	if err := models.DB.Table("diary-data").Where(&diaryInput, user).Find(&diary, no).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	icon := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'favicon'").Find(&Ic).Error
	if icon != nil {
		log.Print(icon.Error())
	}
	result := models.DB.Table("img-asset").Select("`img-asset`.`id`, `img-asset`.`name`, `img-asset`.`img`, `img-path`.`path`").Joins("INNER JOIN `img-path` ON `img-asset`.`path` =`img-path`.`id`").Where("`img-asset`.`name` = 'Fathanah'").Find(&Lg).Error
	if result != nil {
		log.Print(result.Error())
	}
	diaryInput.User = user
	response.Data = diary
	response.Logo = Lg
	response.Icon = Ic
	helper.ResponseJSON(w, http.StatusOK, response)
}

func UpdateDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errr := r.ParseForm()
	if errr != nil {
		panic(errr)
	}
	user, err := strconv.ParseInt(vars["user"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}
	//mengambil user dari parameter
	no, err := strconv.ParseInt(vars["no"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	// memasukan parameter no ke database
	now := time.Now()
	diaryInput := models.Diary{User: user, No: no, Time: now.String()}
	subject := r.Form.Get("subject")
	body := r.Form.Get("body")
	diaryInput.Subject = subject
	diaryInput.Body = body

	// input ke database
	if err := models.DB.Table("diary-data").Where("no = ?", no).Updates(&diaryInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	diaryInput.No = no

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func DeleteDiary(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//mengambil user dari parameter
	no, err := strconv.ParseInt(vars["no"], 10, 64)
	if err != nil {
		helper.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	var diary models.Diary

	// input ke database
	if err := models.DB.Table("diary-data").Where("no = ?", no).Delete(&diary).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
