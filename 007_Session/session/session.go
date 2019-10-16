package session

import "time"

// ISession interface
type ISession interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

// Session : implements the interface
type Session struct {
	ISession
	sid          string                      // unique session id
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // session value stored inside
}

// Set : set session value
func (sess *Session) Set(key, value interface{}) error {
	sess.value[key] = value
	pder.SessionUpdate(sess.sid)
	return nil
}

// Get : get session value
func (sess *Session) Get(key interface{}) interface{} {
	pder.SessionUpdate(sess.sid)
	if v, ok := sess.value[key]; ok {
		return v
	}
	return nil
}

// Delete : delete session value
func (sess *Session) Delete(key interface{}) error {
	delete(sess.value, key)
	pder.SessionUpdate(sess.sid)
	return nil
}

// SessionID : returns current sessionID
func (sess *Session) SessionID() string {
	return sess.sid
}
