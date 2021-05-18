package session

import (
	"fdk-extension-golang/pkg/models"
	"fdk-extension-golang/utils"
	"fmt"

	"github.com/google/uuid"
)

//Session holds session object properties
type Session struct {
	ID           string      `json:"id"`
	CompanyID    string      `json:"company_id"`
	State        string      `json:"state"`
	Scope        []string    `json:"scope"`
	Expires      string      `json:"expires"`
	ExpiresIn    int         `json:"expires_in"`
	AccessMode   string      `json:"access_mode"`
	AccessToken  string      `json:"access_token"`
	CurrentUser  interface{} `json:"current_user"`
	RefreshToken string      `json:"refresh_token"`
	IsNew        bool        `json:"is_new"`
}

//New returns new Session instance
func New(id string, isNew bool) *Session {
	return &Session{
		ID:    id,
		IsNew: isNew,
	}
}

//CloneSession return session copy
func (s *Session) CloneSession(id string, isNew bool) Session {
	newSession := New(id, isNew)
	newSession.CompanyID = s.CompanyID
	newSession.State = s.State
	newSession.Scope = s.Scope
	newSession.Expires = s.Expires
	newSession.ExpiresIn = s.ExpiresIn
	newSession.AccessMode = s.AccessMode
	newSession.AccessToken = s.AccessToken
	newSession.CurrentUser = s.CurrentUser
	newSession.RefreshToken = s.RefreshToken
	return *newSession
}

//GenerateSessionID return session id
func GenerateSessionID(isOnline bool, options models.Option) (string, error) {
	if isOnline == true {
		return uuid.New().String(), nil
	}
	hash, err := utils.GetHash(fmt.Sprintf("%s:%s", options.Cluster, options.CompanyID))
	if err != nil {
		return "", err
	}
	return hash, nil
}
