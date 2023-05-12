/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:32:29 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package middlewares

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/eulbyvan/idk/app/go-user-management/pkg"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Extract the token from the header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		// Parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Provide your own secret key or verification logic here
			// You can retrieve the secret key from a config file or environment variable
			// For example, you can use the same key to sign and verify the token
			secretKey := pkg.GetEnv("SECRET_KEY")
			return []byte(secretKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set the user information in the context for further processing
		// For example, you can extract the user ID from the token and set it in the context
		// The user information can be retrieved in the handler functions using c.GetString("userID")
		claims, _ := token.Claims.(jwt.MapClaims)
		userID := claims["userID"].(string)
		c.Set("userID", userID)

		// Continue to the next middleware or handler
		c.Next()
	}
}
