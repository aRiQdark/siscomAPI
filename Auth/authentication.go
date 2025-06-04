package auth

import (
	"encoding/json"
	models "gin-gonic-gorm/model"
	"log"
	"net/http"
	"golang.org/x/crypto/bcrypt"
)

func login(w *http.ResponseWriter, r *http.Request) {

}

func register(w *http.ResponseWriter, r *http.Request) {
var memberinput models.MemberInput
decode:= json.NewDecoder(r.Body)
if err := decode.Decode(&memberinput);err!= nil{
	log.Fatal("error register")
}
defer r.Body.Close()

bycriptpassword,_ :=bcrypt.GenerateFromPassword([]byte(memberinput.Password),bcrypt.DefaultCost)
memberinput.Password = string(bycriptpassword)

}

func logout(w *http.ResponseWriter, r *http.Request) {

}
