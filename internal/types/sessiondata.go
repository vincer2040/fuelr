package types

import "github.com/gorilla/sessions"

type SessionData struct {
    Authenticated bool
    UserID int64
}

type SessionMissingAuthenticated struct {}
type SessionMissingUserID struct {}

func SessionDataFromSession(session *sessions.Session) (*SessionData, error)  {
    authenticated := session.Values["authenticated"]
    if authenticated == nil {
        return nil, &SessionMissingAuthenticated{}
    }
    userID := session.Values["userID"]
    if userID == nil {
        return nil, &SessionMissingUserID{}
    }
    return &SessionData{
        Authenticated: authenticated.(bool),
        UserID: userID.(int64),
    }, nil
}

func (sessionMissingAuthenticated *SessionMissingAuthenticated) Error() string {
    return "missing authenticated"
}

func (sessionMissingUserID *SessionMissingUserID) Error() string {
    return "missing userID"
}
