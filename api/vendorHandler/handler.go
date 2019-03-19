package vendorHandler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func VendorInit(r *mux.Router) {
	r.HandleFunc("/vendor", list).Methods("GET")
	r.HandleFunc("/vendor", add).Methods("POST")
	r.HandleFunc("/vendor", update).Methods("PUT")
	r.HandleFunc("/vendor/:vendorid", delete).Methods("DELETE")
}

func list(w http.ResponseWriter, r *http.Request) {}

func add(w http.ResponseWriter, r *http.Request) {}

func update(w http.ResponseWriter, r *http.Request) {}

func delete(w http.ResponseWriter, r *http.Request) {}
