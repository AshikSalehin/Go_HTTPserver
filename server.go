package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// BasicAuth middleware to protect private endpoints
func basicAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != "admin" || pass != "password" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized\n"))
			return
		}
		next.ServeHTTP(w, r) // call next handler if authentication passes
	}
}

// Public GET endpoint that returns a JSON response
func publicJSONHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to the public JSON endpoint!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Public HTML form endpoint (GET)
func publicHTMLFormHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("form.html"))
	tmpl.Execute(w, nil)
}

// Public HTML form endpoint (POST) - receiving form data
func publicHTMLFormPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		message := r.FormValue("message")

		// Respond with a simple message back to the client
		fmt.Fprintf(w, "Received: Name - %s, Message - %s", name, message)
	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}

// Private GET endpoint that returns a protected JSON response
func privateJSONHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Welcome to the private JSON endpoint!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Private POST endpoint that accepts JSON and responds with JSON
func privatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data map[string]string
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response := map[string]string{
			"status":  "success",
			"message": "Received your data",
			"data":    data["data"],
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Public endpoints
	http.HandleFunc("/public/json", publicJSONHandler)
	http.HandleFunc("/public/form", publicHTMLFormHandler)
	http.HandleFunc("/public/form-post", publicHTMLFormPostHandler)

	// Private endpoints (protected by basic auth middleware)
	http.HandleFunc("/private/json", basicAuthMiddleware(privateJSONHandler))
	http.HandleFunc("/private/post", basicAuthMiddleware(privatePostHandler))

	fmt.Println("Server running on http://localhost:8080...")
	http.ListenAndServe(":8080", nil)
}
