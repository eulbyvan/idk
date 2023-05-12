/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:03:00 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package domain

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}