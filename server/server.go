/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Sat May 13 2023 12:16:25 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package server

import (
	"log"

	"github.com/eulbyvan/idk/app/go-user-management/api"
	"github.com/eulbyvan/idk/app/go-user-management/config"
)

func Run() error {
	// Initialize the database connection
	db, err := config.InitDB()
	if err != nil {
		return err
	}
	defer config.CloseDB(db)

	// Create a new router
	router := api.SetupRouter(db)

	// Start the server
	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		return err
	}

	return nil
}
