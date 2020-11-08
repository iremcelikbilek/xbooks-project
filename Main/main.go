package main

import (
	"fmt"
	"net/http"
	"os"

	Login "../Modules/Auth/Login"
	Signup "../Modules/Auth/SignUp"
	AddBook "../Modules/Books"
	fb "../Modules/Firebase"
	UpdatePerson "../Modules/UpdateProfile"
	UserData "../Modules/UserData"
	Util "../Modules/Utils"
)

func main() {
	environmentSet()
	go fb.ConnectFirebase()
	createServer()
}

func createServer() {
	go http.HandleFunc("/signup", Signup.HandleSignUp)
	go http.HandleFunc("/login", Login.HandleLogin)
	go http.HandleFunc("/books", AddBook.BookAddHandler)
	go http.HandleFunc("/book", AddBook.GetBooks)
	go http.HandleFunc("/updatePerson", UpdatePerson.UpdateProfileHandler)
	go http.HandleFunc("/userData", UserData.UserDataHandler)
	go http.HandleFunc("/userDataSave", UserData.UserDataSaveHandler)
	go http.HandleFunc("/", handleHome)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		print(err)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	Util.EnableCors(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	_, mail := Util.CheckToken(r)
	fmt.Fprintf(w, mail)
}

func environmentSet() {
	os.Setenv(
		"FIREBASECREDENTIONAL",
		`{  
			"type": "service_account",
			"project_id": "xbooks-hackathon",
			"private_key_id": "860112d40f6ca8647c3657383de629e9d6371997",
			"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQCR2WfpezM7cjQr\nQSCno+88cTvozJi8RGuSwApqOHEy9e2bugbUD0pkTTf2ib3DMRMIEoSr0d7kn0aN\nnsspl3sKhITRSsLoAOEqnstxXSq/havDYn1dbx2n0363tafwdlvXbOSBZwOdXApY\n2gtlKaBX3EYd4ytljmJv+a24FBvJ1cjOpOCw47jvmF4XcUt/lq+r3BADRPA3kG7w\nFAF+3AzNqVMAFdMzW9ZLCK+t5QLLoBYibFR+me9+OpVyssDYtBRArshNgL/VOtky\nr4MjruSWTVboz8mAh5w4+fy0zxJxcpmLktCkrrQ+39XS4luo8V7/3QKwb4H+8ybE\nTxwRkiTPAgMBAAECggEAASdNXsJeWwesiUm1lBMu2YqHum/2e/QKlG7iZjQAvodw\nE/2PqkOzXJcObJZq+CKTj5/IiYIFF50kNgszI3wS5A1nmNl/hc2OpQNRK+9cXtwU\nH10luAR19pWB+Lqovl/L8CiTUhavfPd1DtOCqDqBDdDXLnX/1N4gbIRMBHBcRpIK\n5a19SrEScr7alb2gkB6Mu/xCOLri2AA2oc154qZfKTNTrHhV7xQVRGwwiEro6Y1g\nl8AfN4VEXgtFs/77l3cjUSjMGEafi/v9jmf5Hhe9nclB7RsansoDA/uQaf+7m0OL\n+N5rKY/wn2X1fdEAQbxObipTc+olzg63gWS4av5bCQKBgQDDeew5cwRbOjoipHLG\n7sR9LPsx6EPzk33SjrpIod6frZk5pXA3m7Cs8R8FKeSCirKijCgFTwULUTevdi8U\nFVnVQIRJ8P4cUEtuQKunG7AirvVUGzsEEdXHdOXCkyZ/QewnYovAkkHFBUO19Lu3\nz3GAegGSkxcyo+PvLboddG7PgwKBgQC/AeLQT07qjrIHynkEcaZ6VuBvpOkpT8D4\nkLXnjlWYzITiTaLwVk1Q8SWdG2hyMkwK0ayzdSVTzly/Km/LKbivykaFEzl3Q0/J\n5+A0kdxUKhTRungYn5tfjUzUHKWkTw7486i22hJL6NGE5aU0XERbFGH2/hkQ3AlX\neNZIA/qnxQKBgBzkraUEOjWd0rACLLD44/Q4rNyAn2Kmf2ebDy/jNTvW9hoOORsP\nGTtG3LXvrkYZaDQckWHPf0hf9eIqjuTiDwg1ZBhl1bmrqKqgRn8J2awWvk5zQ/Lj\nC/1Saw+qnnsa7GQZ7dxGAhC0KYEArqqJsBY0cn0O3tPuY6eb07G83xAvAoGASMWm\nLalZrsHpEFDS5J+MRuYcokyZTNUG92zi8TLoZ8vcAQCFSL4IvQLzSrUriP6ivCiA\nfOrv6ssfyPGZIDVoGQme7oqRhV/O4WUHpd/AqwxRXEqIARGmN7R1BgdYEI7SbIUg\nzXGuMN+mG9UGqTlSDcVpQZoz4JmEsZ/nPA7QC+UCgYBgTh+ipbN4xqno10+RsSmR\nnUlwmbEkwLIR73F5WZjFbK/6AHW8CfbkwUaGgK+BhuXpiCqYq+wFlgCtz8PiMqjo\npVUKVyX1YgeIGF223puPMCLo9kvCis9WKr3Y7hSTyrHwftgxx4ARbN4zOVIopMg5\nb1SVDOfosOOiuNatkQEzTg==\n-----END PRIVATE KEY-----\n",
			"client_email": "firebase-adminsdk-6iks5@xbooks-hackathon.iam.gserviceaccount.com",
			"client_id": "105969149645677953015",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-6iks5%40xbooks-hackathon.iam.gserviceaccount.com"
		  }`,
	)

	os.Setenv(
		"SENDINBLUEPASS",
		"xsmtpsib-b4250cc72885a4001608f766a1e22696364c22162fd12754fc7b95d42bd37c8f-FrS8WC2XhkZP4LfD",
	)
}
