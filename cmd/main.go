/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:11:18 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"log"
	"net/http"

	"github.com/eulbyvan/idk/app/go-user-management/api"
)

func main() {
	router := api.SetupRouter()

	// Run the server
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}