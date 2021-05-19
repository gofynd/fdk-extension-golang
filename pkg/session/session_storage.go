package session

import (
	"encoding/json"
	"fdk-extension-golang/pkg/extension"
	"time"
)

//Storage ...
type Storage struct {
	extension *extension.Extension
}

//NewSessionStorage returns new session storage instance
func NewSessionStorage(ext *extension.Extension) *Storage {
	return &Storage{ext}
}

//SaveSession saves session to redis or memory
func (s *Storage) SaveSession(session *Session) error {
	currentTime := time.Now().UTC()
	var err error
	sessionJSONBytes, err := json.Marshal(&session)
	if err != nil {
		return err
	}
	if session.Expires != "" {
		parsedSessionExpiry, err := time.Parse(time.RFC3339, session.Expires)
		if err != nil {
			return err
		}
		if (parsedSessionExpiry != time.Time{}) && (parsedSessionExpiry.After(currentTime)) {
			ttl := parsedSessionExpiry.Sub(currentTime)
			_, err = s.extension.Storage.RedisStorer.Setex(session.ID, string(sessionJSONBytes), ttl)
			if err != nil {
				return err
			}
			return nil
		}
	}
	_, err = s.extension.Storage.RedisStorer.Set(session.ID, string(sessionJSONBytes))
	if err != nil {
		return err
	}
	return nil
}

//GetSession gets session from redis or memory
func (s *Storage) GetSession(sessionID string) (*Session, error) {
	var (
		session, cloneSession Session
		err                   error
	)
	sessionJSONStr, err := s.extension.Storage.RedisStorer.Get(sessionID)
	if err != nil {
		return &Session{}, err
	}
	err = json.Unmarshal([]byte(sessionJSONStr), &session)
	if err != nil {
		return &Session{}, err
	}
	cloneSession = session.CloneSession(sessionID, false)
	return &cloneSession, nil
}

//DeleteSession deletes session from redis or memory
func (s *Storage) DeleteSession(sessionID string) error {
	_, err := s.extension.Storage.RedisStorer.Del(sessionID)
	if err != nil {
		return err
	}
	return nil
}
