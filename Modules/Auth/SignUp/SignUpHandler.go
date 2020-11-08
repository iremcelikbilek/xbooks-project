package SignUp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	fb "../../Firebase"
	util "../../Utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

var JWT_Token = []byte("XBOOKS_JWT_TOKEN")

func HandleSignUp(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	var response util.GeneralResponseModel
	var signupData SignUpModel

	err := json.NewDecoder(r.Body).Decode(&signupData)
	fmt.Println(signupData)

	if err != nil {
		response = util.GeneralResponseModel{
			true, "Gelen veriler hatalı", nil,
		}
		w.Write(response.ToJson())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !util.IsEmailValid(signupData.PersonEmail) {
		response = util.GeneralResponseModel{
			true, "eMail geçersiz", nil,
		}
		w.Write(response.ToJson())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(signupData.PersonName) < 2 || len(signupData.PersonLastName) < 2 {
		response = util.GeneralResponseModel{
			true, "İsim veya Soyisim en az 2 karakterli olmalıdır.", nil,
		}
		w.Write(response.ToJson())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(signupData.Password) < 8 {
		response = util.GeneralResponseModel{
			true, "Parola en az 8 karakterli olmalıdır.", nil,
		}
		w.Write(response.ToJson())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fetchedData := fb.GetFilteredData("/persons", "personEmail", signupData.PersonEmail)
	var result SignUpModel
	mapstructure.Decode(fetchedData, &result)

	if result.PersonEmail != "" {
		response = util.GeneralResponseModel{
			true, "eMail zaten kullanımda", nil,
		}
		w.Write(response.ToJson())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nowDate, _ := time.Now().MarshalText()
	dbData := SignUpDbModel{
		SignUpDate:     string(nowDate),
		SignInDate:     string(nowDate),
		PersonName:     signupData.PersonName,
		PersonLastName: signupData.PersonLastName,
		PersonEmail:    signupData.PersonEmail,
		Password:       util.PasswordHasher(signupData.Password),
	}

	if fbError := fb.PushData("/persons", dbData); fbError != nil {
		response = util.GeneralResponseModel{
			true, fbError.Error(), nil,
		}
		w.Write(response.ToJson())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(1000 * time.Hour)
	claims := &Claims{
		Username: signupData.PersonEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_Token)
	if err != nil {
		response = util.GeneralResponseModel{
			true, "Bir Sorun Oluştu", nil,
		}
		w.Write(response.ToJson())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tokenData := LoginResponseData{
		Token:   tokenString,
		Expires: expirationTime.String(),
	}

	response = util.GeneralResponseModel{
		false, "Kayıt Başarılı", tokenData,
	}

	var data = UserDefaultData{
		UserMail:       signupData.PersonEmail,
		UserTotalPoint: 1000,
	}
	var hata = fb.PushData("/userData", data)
	fmt.Println(hata)
	w.Write(response.ToJson())
}

type UserDefaultData struct {
	UserMail       string `json:"userMail"`
	UserTotalPoint int    `json:"userTotalPoint"`
}
