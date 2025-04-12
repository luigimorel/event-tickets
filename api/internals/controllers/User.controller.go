package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/morelmiles/go-events/internals/errors"
	"github.com/morelmiles/go-events/internals/helpers"
	"github.com/morelmiles/go-events/internals/models"
	"github.com/morelmiles/go-events/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to fetch users")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}
	helpers.JSON(w, http.StatusOK, &users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		apiErr := errors.NewBadRequestError("Invalid user ID format")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		apiErr := errors.NewNotFoundError("User not found")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}
	helpers.JSON(w, http.StatusOK, user)
}

func GetAllEventsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		apiErr := errors.NewBadRequestError("Invalid user ID format")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	var user models.User
	if err := database.DB.Preload("Events").First(&user, id).Error; err != nil {
		apiErr := errors.NewNotFoundError("User not found")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}
	helpers.JSON(w, http.StatusOK, user.Events)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := helpers.DecodeJSONBody(w, r, &user); err != nil {
		apiErr := errors.NewValidationError("Invalid request body")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := user.BeforeSave(database.DB); err != nil {
		apiErr := errors.NewInternalServerError("Failed to process user data")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to create user")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	helpers.JSON(w, http.StatusCreated, user)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		apiErr := errors.NewBadRequestError("Invalid user ID format")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		apiErr := errors.NewNotFoundError("User not found")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := helpers.DecodeJSONBody(w, r, &user); err != nil {
		apiErr := errors.NewValidationError("Invalid request body")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := database.DB.Save(&user).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to update user")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	helpers.JSON(w, http.StatusOK, user)
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		apiErr := errors.NewBadRequestError("Invalid user ID format")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		apiErr := errors.NewNotFoundError("User not found")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		apiErr := errors.NewDatabaseError("Failed to delete user")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	helpers.JSON(w, http.StatusOK, user)
}

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		apiErr := errors.NewBadRequestError("Failed to read request body")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	user := models.User{}
	if err = json.Unmarshal(body, &user); err != nil {
		apiErr := errors.NewValidationError("Invalid request body format")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	user.Prepare()

	if err = user.Validate("login"); err != nil {
		apiErr := errors.NewValidationError(err.Error())
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		apiErr := errors.NewUnauthorizedError("Invalid credentials")
		helpers.ERROR(w, apiErr.StatusCode, apiErr)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "jwt",
		Value: token,
		Path:  "/",
	})

	helpers.JSON(w, http.StatusOK, token)
}

func SignIn(email, password string) (string, error) {
	var err error

	user := models.User{}
	err = database.DB.Debug().Model(models.User{}).Where("email=?", email).Take(&user).Error

	if err != nil {
		return " ", err
	}

	err = models.VerifyPassword(password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return helpers.CreateToken(user.ID)
}
