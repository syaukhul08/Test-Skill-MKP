package controllers

import (
	"GO-TEST-MKP/configs"
	"GO-TEST-MKP/helpers"
	"GO-TEST-MKP/models"
	"encoding/json"
	"net/http"
)

func Terminal(w http.ResponseWriter, r *http.Request) {

	var terminal models.Terminal

	user := r.Context().Value("userinfo").(*helpers.MyCustomClaims)
	userResponse := &models.Transaction{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Terminals: []models.Terminal{},
	}

	err := json.NewDecoder(r.Body).Decode(&terminal)
	if err != nil {
		helpers.Response(w, http.StatusBadRequest, "Invalid JSON data", nil)
		return
	}
	
	if err := configs.DB.Create(&terminal).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Terminal added", userResponse)
}