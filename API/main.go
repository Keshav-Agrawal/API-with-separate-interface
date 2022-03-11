package main

import (
	"fmt"

	"github.com/Keshav-Agrawal/mongoapi/datasource/mongo"

	"log"
	"net/http"

	"github.com/Keshav-Agrawal/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	ds := mongo.NewDs()
	r := router.Router(ds)
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000 ...")
}
