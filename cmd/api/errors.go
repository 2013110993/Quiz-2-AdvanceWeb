// Filename : cmd/api/errors.go

package main

import (
	"fmt"
	"net/http"
)

// Log errors
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// Error message in JSON-formatted
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	// create the response
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}

}

// Server error message
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//log the error
	app.logError(r, err)
	//A message with error
	message := "the server encountered an error and could not proceed "
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Not found response
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	//prepare a message with error
	message := "resources could not be found."
	app.errorResponse(w, r, http.StatusNotFound, message)
}

// Not Allowed response
func (app *application) MethodNotAllowedReponse(w http.ResponseWriter, r *http.Request) {
	//prepare a message with error
	message := fmt.Sprintf(" method: %s is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// Validation errors
func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
