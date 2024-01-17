package models

// API user credentials
// Used to allow users to sign into the API
//
// swagger:model user
type User struct {
	// User's login name
	//
	// required:true
	Username string `json:"username"`
	// User's email
	//
	// required:true
	Email string `json:"email"`
	// User's password
	//
	// required:true
	Password string `json:"password"`
}
