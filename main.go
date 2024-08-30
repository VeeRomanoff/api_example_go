package main

import (
	"SEMI_TRASH_API/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	// тут проосто константа как в @RequestMapping
	apiPrefix string = "/api/v1"
)

var (
	port string
	// ресурсы которые мы определили, как неизменные пути
	bookResourcePrefix      string = apiPrefix + "/book"  // api/v1/book
	manyBooksResourcePrefix string = apiPrefix + "/books" // api/v1/books
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load .env file")
	}

	port = os.Getenv("app_port")
}

func main() {
	log.Printf("starting rest api application on port: %s", port)
	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResource(router, manyBooksResourcePrefix)

	log.Println("Router initilized successfully. Ready to go.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
