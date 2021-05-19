package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//SESSIONCOOKIENAME ...
const SESSIONCOOKIENAME = "ext_session"

//Option holds options object properties
type Option struct {
	CompanyID string `json:"company_id"`
	Cluster   string `json:"cluster"`
}

//Callback ...
type Callback func(map[string]interface{}) string

//ExtCallback ...
type ExtCallback struct {
	// Setup     Callback
	Install   Callback
	Auth      Callback
	Uninstall Callback
}

//User model
type User struct {
	ID        string `json:"_id"`
	UID       string `json:"uid"`
	UserName  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// PhoneNumber        []PhoneNumber `json:"phoneNumbers"`
	// Emails             []Email       `json:"emails"`
	Gender             string        `json:"gender"`
	AccountType        string        `json:"accountType"`
	Image              string        `json:"image"`
	Roles              []interface{} `json:"roles"`
	Active             bool          `json:"active"`
	ProfilePicURL      string        `json:"profilePicUrl"`
	CreatedAt          time.Time     `json:"createdAt"`
	HasOldPasswordHash bool          `json:"hasOldPasswordHash"`
	UpdatedAt          time.Time     `json:"updatedAt"`
	Hash               string        `json:"hash"`
	Debug              Debug         `json:"debug"`
}

//Debug model
type Debug struct {
	Source   string `json:"source"`
	Platform string `json:"platform"`
}

// Application is Slingshot Application
type Application struct {
	ID          primitive.ObjectID `json:"_id"`
	Description string             `json:"description"`
	CacheTTL    int                `json:"cache_ttl"`
	Name        string             `json:"name"`
	Owner       string             `json:"owner"`
	Token       string             `json:"token"`
	Secret      string             `json:"secret"`
	CompanyID   int32              `json:"company_id"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`

	Domain struct {
		Verified bool   `json:"verified"`
		Name     string `json:"name"`
	} `json:"domain"`

	Website struct {
		Enabled  bool   `json:"enabled"`
		Basepath string `json:"basepath"`
	} `json:"website"`

	Cors struct {
		Domains []string `json:"domains"`
	} `json:"cors"`

	Meta []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"meta"`
}
