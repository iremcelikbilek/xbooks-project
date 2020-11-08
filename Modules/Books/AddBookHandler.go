package Books

import (
	"encoding/json"
	"net/http"

	fb "../Firebase"
	util "../Utils"
)

func BookAddHandler(w http.ResponseWriter, r *http.Request) {

	var response util.GeneralResponseModel
	var addBookData AddBookDataModel

	if err := json.NewDecoder(r.Body).Decode(&addBookData); err != nil {
		writeError("Gelen veriler hatalı", w)
		return
	}

	dbData := AddBookDataModel{
		BookName:    addBookData.BookName,
		Author:      addBookData.Author,
		Explanation: addBookData.Explanation,
		BookURL:     addBookData.BookURL,
		BookCost:    addBookData.BookCost,
	}

	if saveErr := fb.PushData("/books", dbData); saveErr != nil {
		writeError("Konum bilgisi hatalı", w)
		return
	}

	response = util.GeneralResponseModel{
		false, "Kitap başarıyla eklendi.", nil,
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
