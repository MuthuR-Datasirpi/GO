package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	serve()
// 	r := mux.NewRouter()

// 	r.HandleFunc("/", serveHTTP)

// 	server := http.Server{
// 		Addr:    ":8010",
// 		Handler: r,
// 	}
// 	server.ListenAndServe()

// }

// func serve() {
// 	fmt.Println("Hello the program is start")
// }

// func serveHTTP(w http.ResponseWriter, r *http.Request) {
// 	// w.Write([]byte("<h1>Welcome to the program </h1>"))   //! everything is from web coming in byte so we use []byte()

// 	log.Fatal(fmt.Fprintf(w, "<h1>welcome muthu</h1>"))
// }
