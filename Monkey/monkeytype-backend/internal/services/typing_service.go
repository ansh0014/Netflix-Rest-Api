package services

import (
	"errors"
	"time"
)

type TypingSession struct {
	UserID    string
	Text      string
	Timestamp time.Time
}

type TypingService struct {
	sessions map[string]*TypingSession
}

func NewTypingService() *TypingService {
	return &TypingService{
		sessions: make(map[string]*TypingSession),
	}
}

func (ts *TypingService) StartTypingSession(userID string, text string) (*TypingSession, error) {
	if userID == "" {
		return nil, errors.New("userID cannot be empty")
	}
	session := &TypingSession{
		UserID:    userID,
		Text:      text,
		Timestamp: time.Now(),
	}
	ts.sessions[userID] = session
	return session, nil
}

func (ts *TypingService) EndTypingSession(userID string) (*TypingSession, error) {
	session, exists := ts.sessions[userID]
	if !exists {
		return nil, errors.New("no typing session found for user")
	}
	delete(ts.sessions, userID)
	return session, nil
}