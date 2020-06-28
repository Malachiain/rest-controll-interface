package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// HomeHandler  handle root path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("called root")
	fmt.Fprintf(w, "OK")
}

// Angle an angle of rotation
type Angle struct {
	Vertical   int
	Horizontal int
}

// RotateToAngle rotate to given angle given it Angle.To
func RotateToAngle(w http.ResponseWriter, r *http.Request) {
	var a Angle
	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		fmt.Fprint(w, "Bad request")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if a.Horizontal > 180 || a.Horizontal < 0 {
		http.Error(w, "Bad request, horizontal angle must be greater than  0 and less than  or equal to 180", http.StatusBadRequest)
		return
	} else if a.Vertical > 180 || a.Vertical < 0 {
		http.Error(w, "Bad request, vertical angle must be greater than  0 and less than  or equal to 180", http.StatusBadRequest)
		return
	} else {
		log.Printf("rotating to horizontal: %+v vertical: %+v", a.Horizontal, a.Vertical)
		// Do something with the Person struct...
		fmt.Fprintf(w, "New angle: %+v %+v \n", a.Vertical, a.Horizontal)
	}
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/rotate", RotateToAngle)
	http.Handle("/", r)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("starting server")

	http.ListenAndServe(":8080", nil)
}
