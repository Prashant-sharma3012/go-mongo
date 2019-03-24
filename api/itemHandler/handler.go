package itemHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tryTwo/utils"
)

func LoadItemHandler(r *mux.Router) {
	r.HandleFunc("/item", list).Methods("GET")
	r.HandleFunc("/item", add).Methods("POST")
	r.HandleFunc("/item", update).Methods("PUT")
	r.HandleFunc("/item", delete).Methods("DELETE")
}

func list(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.ParseInt(r.FormValue("limit"), 10, 64)
	skip, _ := strconv.ParseInt(r.FormValue("skip"), 10, 64)

	items, err := ItemList(skip, limit)

	if err != nil {
		utils.Error(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	jsonRes, _ := json.Marshal(items)
	w.Write(jsonRes)
}

func add(w http.ResponseWriter, r *http.Request) {
	utils.Log("Saving an Item")

	item := ItemFromJson(r.Body)
	_, err := item.Save()

	if err != nil {
		utils.Error(err.Error())
		w.Write([]byte("Something went wrong"))
	} else {
		utils.Log("Item Saved Successfully")
		w.Write([]byte("Item Added Successfully"))
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	utils.Log("Updating an Item")
	item := ItemFromJson(r.Body) // this is just to pull itemID
	_, err := item.Update()

	if err != nil {
		utils.Error(err.Error())
		w.Write([]byte("Something went wrong"))
	} else {
		utils.Log("Item Updated Successfully")
		w.Write([]byte("Item Updated Successfully"))
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	utils.Log("Updating an Item")

	item := ItemFromJson(r.Body)
	_, err := item.Delete()

	if err != nil {
		utils.Error(err.Error())
		w.Write([]byte("Something went wrong"))
	} else {
		utils.Log("Item Deleted Successfully")
		w.Write([]byte("Item Deleted Successfully"))
	}
}
