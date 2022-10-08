// File: cmd/api/random.go

package main

import "net/http"

func (app *application) randomHandler(w http.ResponseWriter, r *http.Request) {
	//From helpers.go
	id, err := app.readIDParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Extract random value based on :id
	random_value := app.generateRandomString(int(id))
	err = app.writeJSON(w, http.StatusOK, envelope{"random_value": random_value}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
