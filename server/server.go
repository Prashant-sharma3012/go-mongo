package server

import (
	"net/http"
	"os"

	"github.com/tryTwo/db"
	"github.com/tryTwo/utils"

	"github.com/gorilla/mux"
	"github.com/tryTwo/api/itemHandler"
	"github.com/tryTwo/api/vendorHandler"
)

type Server struct {
	Routes *mux.Router
}

func BootstrapServer() *Server {
	server := &Server{}
	server.Routes = mux.NewRouter().StrictSlash(true)

	err := db.Connect()
	if err != nil {
		utils.Error("DB Connect error, Exiting")
		os.Exit(1)
	}

	initRoutes(server)
	return server
}

func initRoutes(s *Server) {
	s.Routes.HandleFunc("/", ping).Methods("GET")
	itemHandler.LoadItemHandler(s.Routes)
	vendorHandler.VendorInit(s.Routes)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is up and running"))
}
