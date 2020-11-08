package UpdateProfile

import "github.com/dgrijalva/jwt-go"

type UpdateProfileDataModel struct {
	PersonName     string `json:"personName"`
	PersonLastName string `json:"personLastName"`
	Password       string `json:"password"`
}

type Claims struct {
	Username string `json:"personEmail"`
	jwt.StandardClaims
}

type UpdateDbModel struct {
	UpdateDate     string `json:"signUpDate"`
	PersonName     string `json:"personName"`
	PersonLastName string `json:"personLastName"`
	PersonEmail    string `json:"personEmail"`
	Password       string `json:"password"`
}

type UpdateResponseData struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}
