package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func serverHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "welcome to my website")
// }
// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Home page")
// }

// func main() {
// 	fmt.Println("Hello world")

// 	// router
// 	router := http.NewServeMux() //!

// 	// handler
// 	router.HandleFunc("/", home)
// 	router.HandleFunc("/login", serverHTTP)

// 	// server
// 	server := http.Server{
// 		Addr:    ":8080",
// 		Handler: router,
// 	}

// 	server.ListenAndServe()

// }
