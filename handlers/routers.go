package handlers

import (

	"fmt"
	"github.com/digvijay17july/go-server-server/models"
	"github.com/digvijay17july/go-server-server/utils"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"

)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// App initialize with predefined configuration
func (a *App) Initialize(config *utils.Config) {
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		config.DB.Username,
		config.DB.Password,
		config.DB.Host,
		config.DB.PortNo,
		config.DB.Name,
		config.DB.Charset,
	)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database %s",err.Error())
	}
	a.DB = DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()

}
// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&models.Person{})
	return db
}
func (a *App) setRouters(){
	fmt.Println("initializing request")
	a.Get("/", a.init)
	a.Get("/person/{uuid}", a.GetUser)
	a.Post("/person", a.CreateUser)
	a.Get("/people", a.GetUsers)
}

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser(a.DB, w, r)
}

func (a *App) GetUsers(w http.ResponseWriter, r *http.Request) {
	GetPeople(a.DB, w, r)
}

func (a *App) GetUser(w http.ResponseWriter, r *http.Request) {
	GetUser(a.DB, w, r)
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App)init(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello")
}