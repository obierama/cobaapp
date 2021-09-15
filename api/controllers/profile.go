package controllers

import (
	"coba/api/model"
	"coba/config/respon"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	respn1 = map[string]interface{}{"status": true, "message": "Succes", "code": 200}
)

func (a *App) Profile(w http.ResponseWriter, r *http.Request) {

	profile := &model.Profile{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//buku.Prepare()

	err = profile.Validate("")
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userCreated, err := profile.SaveProfile(a.DB)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = userCreated
	respon.JSON(w, http.StatusCreated, resp)
	return

}

func (a *App) UpdateProfile(w http.ResponseWriter, r *http.Request) {

	profile := &model.Profile{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &profile)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = profile.Validate("update")
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}
	userCreated, err := profile.Update(profile.id("id"), a.DB)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}
	respn1["data"] = userCreated
	respon.JSON(w, http.StatusCreated, resp)
	return
}
