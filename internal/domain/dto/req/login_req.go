/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Wed May 17 2023 8:12:59 PM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package req

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}