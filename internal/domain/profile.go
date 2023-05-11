/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:03:16 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package domain

type Profile struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Bio    string `json:"bio"`
}
