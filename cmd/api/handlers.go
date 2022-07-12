package main

import (
	"net/http"
)

// Struct for returning responses.
type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// Login Handler
func (app *application) Login(w http.ResponseWriter, r *http.Request) {

	// for accepting and storing request data.
	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// struct to receive incoming credentials and decode incoming requests into the response.
	var creds credentials

	var payload jsonResponse

	err := app.readJSON(w, r, &creds)
	if err != nil {
		payload.Error = true
		payload.Message = "Invalid JSON supplied or JSON missing entirely"
		_ = app.writeJSON(w, http.StatusBadRequest, payload)
	}

	// authenticate
	app.infoLog.Printf("Email: %v - Password: %v", creds.Email, creds.Password)

	// send back a response
	payload.Error = false
	payload.Message = "Signed In"

	// out, err := json.MarshalIndent(payload, "", "\t")
	err = app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorLog.Println(err)
	}

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(out)

}
