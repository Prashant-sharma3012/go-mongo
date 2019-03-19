package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tryTwo/server"
)

func main() {
	fmt.Println("setting up server...")

	srv := server.BootstrapServer()

	fmt.Println("server setup complete, starting ....")

	err := http.ListenAndServe(":8000", srv.Routes)

	if err != nil {
		fmt.Println("Something went wrong")
		fmt.Println(err.Error())
		fmt.Println("######### Exiting #############")
		os.Exit(1)
	}

	fmt.Println("Server is up and listening on 8000")
}
