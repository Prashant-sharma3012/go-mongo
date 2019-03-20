package itemHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ItemInit(r *mux.Router) {
	r.HandleFunc("/item", list).Methods("GET")
	r.HandleFunc("/item", add).Methods("POST")
	r.HandleFunc("/item", update).Methods("PUT")
	r.HandleFunc("/item", delete).Methods("DELETE")
}

func list(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.ParseInt(r.FormValue("limit"), 10, 64)
	skip, _ := strconv.ParseInt(r.FormValue("skip"), 10, 64)

	items := List(skip, limit)

	jsonRes, _ := json.Marshal(items)
	w.Write(jsonRes)
}

func add(w http.ResponseWriter, r *http.Request) {
	item := ItemFromJson(r.Body)
	_, err := item.Save()

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("Something went wrong"))
	} else {
		w.Write([]byte("Item Added Successfully"))
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	item := ItemFromJson(r.Body)
	_, err := item.Update()

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("Something went wrong"))
	} else {
		w.Write([]byte("Item Updated Successfully"))
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	item := ItemFromJson(r.Body)
	_, err := item.Delete()

	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte("Something went wrong"))
	} else {
		w.Write([]byte("Item Deleted Successfully"))
	}
}
