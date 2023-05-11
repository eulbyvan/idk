/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 12:14:50 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Invalid request method. Only GET is allowed.")
		return
	}

	fmt.Fprintf(w, "Hello, world!")
}