package UserData

import (
	"encoding/json"
	"net/http"

	fb "../Firebase"
	util "../Utils"
)

func UserDataHandler(w http.ResponseWriter, r *http.Request) {
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

	fetchedData := fb.GetFilteredData("/userData", "userMail", userMail)
	response = util.GeneralResponseModel{
		true, "Başarılı", fetchedData,
	}
	w.Write(response.ToJson())
}

func UserDataSaveHandler(w http.ResponseWriter, r *http.Request) {
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

	var data interface{}
	json.NewDecoder(r.Body).Decode(&data)

	itemsMap := data.(map[string]interface{})
	itemsMap["userMail"] = userMail

	var response util.GeneralResponseModel
	fb.UpdateFilteredData("/userData", "userMail", userMail, data)

	response = util.GeneralResponseModel{
		true, "", nil,
	}
	w.Write(response.ToJson())
}

type MailAdd struct {
	userMail string
}

func writeError(description string, w http.ResponseWriter) {
	var response util.GeneralResponseModel
	w.WriteHeader(http.StatusBadRequest)
	response = util.GeneralResponseModel{
		true, description, nil,
	}
	w.Write(response.ToJson())
}
