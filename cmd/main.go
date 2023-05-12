/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:11:18 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"log"

	"github.com/eulbyvan/idk/app/go-user-management/server"
)

func main() {
	// Run the server
	if err := server.Run(); err != nil {
		log.Fatalf("Error running the server: %s", err)
	}
}