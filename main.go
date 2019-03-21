package main

import (
	"net/http"
	"os"

	"github.com/tryTwo/utils"

	"github.com/tryTwo/server"
)

func main() {
	utils.Log("setting up server")

	srv := server.BootstrapServer()

	utils.Log("server setup complete, Listening on 8000")

	err := http.ListenAndServe(":8000", srv.Routes)

	if err != nil {
		utils.Error("Something went wrong")
		utils.Error(err.Error())
		utils.Error("######### Exiting #############")
		os.Exit(1)
	}
}
