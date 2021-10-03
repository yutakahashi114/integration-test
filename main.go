package main

import (
	"encoding/json"
	"fmt"
	"integration/openapi"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	db, err := connectDB()
	if err != nil {
		log.Fatalln(err)
		return
	}
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)

	openapi.HandlerFromMux(Controller{db}, mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	http.ListenAndServe(":"+port, mux)
}

type User struct {
	gorm.Model
	Email string `db:"email"`
	Name  string `db:"name"`
}

type Controller struct {
	*gorm.DB
}

func (c Controller) FindUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	err := c.DB.Order("id ASC").Find(&users).Error
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	us := make([]openapi.User, len(users))
	for i, user := range users {
		us[i] = openapi.User{
			Id:        int64(user.ID),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
	}
	render.JSON(w, r, us)
}

func (c Controller) GetUserByID(w http.ResponseWriter, r *http.Request, id int64) {
	user := User{}
	err := c.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, openapi.User{
		Id:        int64(user.ID),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}

func (c Controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := openapi.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	user := &User{
		Name:  u.Name,
		Email: u.Email,
	}
	err = c.DB.Create(&user).Error
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	render.JSON(w, r, openapi.ID{Id: int64(user.ID)})
}

func (c Controller) UpdateUser(w http.ResponseWriter, r *http.Request, id int64) {
	u := openapi.User{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	user := &User{
		Name:  u.Name,
		Email: u.Email,
	}
	err = c.DB.Where("id = ?", id).Updates(&user).Error
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c Controller) DeleteUser(w http.ResponseWriter, r *http.Request, id int64) {
	err := c.DB.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func connectDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
