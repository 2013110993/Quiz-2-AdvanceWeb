// File: cmd/api/info.go

package main

import (
	"net/http"
	"time"

	"fileuploading.federicorosado.net/internal/data"
	"fileuploading.federicorosado.net/internal/validator"
)

// Create a New user
func (app *application) createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Decode destination

	var input struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Email   string `json:"email"`
		Address string `json:"address"`
		School  string `json:"school"`
		Degree  string `json:"degree"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badResquestReponse(w, r, err)
		return
	}
	//New user struct with a copy input values from the input struct
	user := &data.User{
		ID:      1,
		Name:    input.Name,
		Phone:   input.Phone,
		Email:   input.Email,
		Address: input.Address,
		School:  input.School,
		Degree:  input.Degree,
		Version: 1.0,
	}

	// Creates a new instance
	v := validator.New()

	// Check the errors maps if there were any errors validation
	if data.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Display valid input
	err = app.writeJSON(w, http.StatusOK, envelope{"user_data": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// GET User /v1/personalinfo endpoints
func (app *application) showInfoHandler(w http.ResponseWriter, r *http.Request) {
	// sample data
	user := data.User{
		ID:        10,
		CreatedAt: time.Now(),
		Name:      "Federico Rosado",
		Phone:     "630 - 2525",
		Email:     "federicorickyrosado@gmail.com",
		Address:   "Chi Ha Street",
		School:    "UB",
		Degree:    "BINT",
		Version:   1.0,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"user_data": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
