package main

import (

	"log"
	"net/http"
	"github.com/gorilla/mux"
	"main/Routers"
)


// Main function
func main() {

	// Init router
	r := mux.NewRouter()
	Routers.DefineRoutes(r);
	
	// Start API server
    log.Fatal(http.ListenAndServe(":8001", r))
	
}
