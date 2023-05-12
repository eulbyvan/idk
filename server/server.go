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
	"github.com/eulbyvan/idk/app/go-user-management/pkg"
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
	serverAddress := pkg.GetEnv("SERVER_ADDRESS")
	log.Printf("Server is running on address: %s\n", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		return err
	}

	return nil
}
