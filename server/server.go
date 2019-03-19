package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/tryTwo/api/itemHandler"
	"github.com/tryTwo/api/vendorHandler"
	"github.com/tryTwo/db"
)

type Server struct {
	DBConn *db.DB
	Routes *mux.Router
}

func BootstrapServer() *Server {
	server := &Server{}
	server.Routes = mux.NewRouter().StrictSlash(true)

	err := db.Connect()
	if err != nil {
		fmt.Println("DB Connect error, Exiting....")
		os.Exit(1)
	}

	initRoutes(server)
	return server
}

func initRoutes(s *Server) {
	s.Routes.HandleFunc("/", ping).Methods("GET")
	itemHandler.ItemInit(s.Routes)
	vendorHandler.VendorInit(s.Routes)
}

func ping(w http.ResponseWriter, r *http.Request) {}
