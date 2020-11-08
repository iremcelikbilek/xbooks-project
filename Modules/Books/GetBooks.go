package Books

import (
	"net/http"

	fb "../Firebase"
	util "../Utils"

	"github.com/mitchellh/mapstructure"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var response util.GeneralResponseModel
	var bookDatas []AddBookDataModel

	allData := fb.ReadData("/books")

	if allData == nil {
		w.WriteHeader(http.StatusInternalServerError)
		response = util.GeneralResponseModel{
			true, "Bir hata oluştu", nil,
		}
		w.Write(response.ToJson())
		return
	}

	itemsMap := allData.(map[string]interface{})
	for _, data := range itemsMap {
		var bookData AddBookDataModel
		mapstructure.Decode(data, &bookData)
		bookDatas = append(bookDatas, bookData)
	}

	response = util.GeneralResponseModel{
		false, "Veriler başarıyla getirildi", bookDatas,
	}
	w.Write(response.ToJson())
}
