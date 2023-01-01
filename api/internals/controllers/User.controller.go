package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morelmiles/go-events/config"
	"github.com/morelmiles/go-events/internals/helpers"
	"github.com/morelmiles/go-events/internals/models"
	"golang.org/x/crypto/bcrypt"
)

// Compares the password before saving it
func ComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Checks if the user exists
func checkIfUserExists(userId string) bool {
	var user models.User
	config.DB.First(&user, userId)

	return user.ID != 0
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	var users []models.User
	config.DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
}

// GetUserById - Fetches a list of all users.
func GetUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	if !checkIfUserExists(userId) {
		json.NewEncoder(w).Encode("user not found!")
		return
	}

	var user models.User
	config.DB.First(&user, userId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser - Creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	user.BeforeSave()
	err = user.Validate("")
	if err != nil {
		helpers.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// UpdateUserById -  Updates a single user by the ID specified
func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found!")
		return
	}

	var user models.User

	config.DB.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUserById - Updates a single user by the ID specified.
func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userId := mux.Vars(r)["id"]
	if !checkIfUserExists(userId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("user not found!")
		return
	}

	var user models.User
	config.DB.Delete(&user, userId)
	json.NewEncoder(w).Encode(user)
}

func GetAllEventsByUser(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]

	var user models.User
	var event models.Event

	config.DB.Model(&user).Find(event).Where("id = ?", userId)
	json.NewEncoder(w).Encode(user)
}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		helpers.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()

	err = user.Validate("login")
	if err != nil {
		helpers.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := helpers.FormatError(err.Error())
		helpers.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	helpers.JSON(w, http.StatusOK, token)

}

// Sign in

func SignIn(email, password string) (string, error) {
	var err error

	user := models.User{}
	err = config.DB.Debug().Model(models.User{}).Where("email=?", email).Take(&user).Error

	if err != nil {
		return " ", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return helpers.CreateToken(user.ID)
}
