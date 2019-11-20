package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/digvijay17july/go-server-server/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func GetPeople(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	people:=[] models.Person{}
	db.Set("gorm:auto_preload", true).Find(&people)
	RespondJSON(w,http.StatusOK,people)
}
func CreateUser(db *gorm.DB,w http.ResponseWriter, r *http.Request){
	person:= models.Person{}
	decoder := json.NewDecoder(r.Body)
	if err:=decoder.Decode(&person); err!=nil{
		RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer r.Body.Close()
	person.Uuid=createHash(person.Name)
	err :=db.Create(&person)
	if  err.Error!=nil{
		fmt.Println(err.Error.Error())
		RespondError(w, http.StatusBadRequest, err.Error.Error())
		return
	}
	RespondJSON(w, http.StatusCreated,person)
}
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["uuid"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err!=nil{
		fmt.Println(err.Error())
		RespondError(w, http.StatusBadRequest, err.Error())
	}
	person := getUserOr404(db, uint(i), w, r)
	if person == nil {
		return
	}
	RespondJSON(w, http.StatusOK, person)
}
func getUserOr404(db *gorm.DB, id uint, w http.ResponseWriter, r *http.Request) *models.Person {
	person := models.Person{}
	if err := db.Set("gorm:auto_preload", true).First(&person,id).Error; err != nil {
		RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &person
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}