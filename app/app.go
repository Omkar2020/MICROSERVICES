package app

import (
	"log"
	"net/http"
)

func Start() {

	//define routes
	http.HandleFunc("/hellogo", greet)
	http.HandleFunc("/customers", getAllCustomers)
	// http.HandleFunc("/hellogo", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello, Go!")
	// })

	//starting Server
	//http.ListenAndServe("localhost:9000", nil)
	log.Fatal(http.ListenAndServe("localhost:9000", nil))
}
