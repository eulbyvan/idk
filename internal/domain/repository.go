/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Fri May 12 2023 1:05:26 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package domain

type UserRepository interface {
	CreateUser(user *User) error
	GetByEmail(email string) (*User, error)
	// Add other user repository methods as needed
}

type ProfileRepository interface {
	CreateProfile(profile *Profile) error
	// Add other profile repository methods as needed
}
