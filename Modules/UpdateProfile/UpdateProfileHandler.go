package UpdateProfile

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	fb "../Firebase"
	util "../Utils"
	"github.com/mitchellh/mapstructure"
)

func UpdateProfileHandler(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var userMail string
	if isSucessToken, message := util.CheckToken(r); !isSucessToken {
		writeError(message, w)
		return
	} else {
		userMail = message
	}

	var response util.GeneralResponseModel
	var updateProfileDataModel UpdateProfileDataModel

	if err := json.NewDecoder(r.Body).Decode(&updateProfileDataModel); err != nil {
		writeError("Gelen veriler hatalı", w)
		return
	}

	fetchedData := fb.GetFilteredData("/persons", "personEmail", userMail)
	var updateDbModel UpdateDbModel
	mapstructure.Decode(fetchedData, &updateDbModel)

	nowDate, _ := time.Now().MarshalText()

	if updateProfileDataModel.Password == "" {
		updateProfileDataModel.Password = updateDbModel.Password
	} else {
		updateProfileDataModel.Password = util.PasswordHasher(updateProfileDataModel.Password)
	}

	dbData := UpdateDbModel{
		UpdateDate:     string(nowDate),
		PersonName:     updateProfileDataModel.PersonName,
		PersonLastName: updateProfileDataModel.PersonLastName,
		PersonEmail:    userMail,
		Password:       updateProfileDataModel.Password,
	}

	fmt.Println(dbData)

	if err := fb.UpdateFilteredData("/persons", "personEmail", userMail, dbData); err != nil {
		writeError("Güncelleme hatalı", w)
		return
	}

	response = util.GeneralResponseModel{
		false, "Güncelleme tamamlandı", nil,
	}
	w.Write(response.ToJson())
}

func writeError(description string, w http.ResponseWriter) {
	var response util.GeneralResponseModel
	w.WriteHeader(http.StatusBadRequest)
	response = util.GeneralResponseModel{
		true, description, nil,
	}
	w.Write(response.ToJson())
}
