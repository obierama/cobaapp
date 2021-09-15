package controllers

import (
	"coba/api/model"
	"coba/config/respon"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	respn = map[string]interface{}{"status": true, "message": "Succes", "code": 200}
)

func (a *App) Buku(w http.ResponseWriter, r *http.Request) {

	buku := &model.Buku{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &buku)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//buku.Prepare()

	err = buku.Validate("")
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}

	userCreated, err := buku.SaveBuku(a.DB)
	if err != nil {
		respon.ERROR(w, http.StatusBadRequest, err)
		return
	}
	resp["data"] = userCreated
	respon.JSON(w, http.StatusCreated, resp)
	return

}
