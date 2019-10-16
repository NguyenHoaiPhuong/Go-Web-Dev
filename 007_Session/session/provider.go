package session

import (
	"container/list"
	"sync"
	"time"
)

// IProvider interface
type IProvider interface {
	// SessionInit implements the initialization of a session,
	// and returns a new session if it succeeds
	SessionInit(sid string) (ISession, error)

	// SessionRead returns a session represented by the corresponding sid.
	// Creates a new session and returns it if it does not already exist
	SessionRead(sid string) (ISession, error)

	// SessionDestroy given an sid, deletes the corresponding session
	SessionDestroy(sid string) error

	// SessionGC deletes expired session variables according to maxLifeTime
	SessionGC(maxLifeTime int64)
}

var providers = make(map[string]IProvider)

// Register makes a session provider available by the provided name.
// If a Register is called twice with the same name or if the driver is nil,
// it panics.
func Register(name string, provider IProvider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	providers[name] = provider
}

// Provider : implements the interface
type Provider struct {
	IProvider

	lock     sync.Mutex               // lock
	sessions map[string]*list.Element // save in memory
	list     *list.List               // gc
}

// SessionInit : initialize the session
func (pder *Provider) SessionInit(sid string) (ISession, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &Session{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

// SessionRead : returns a session represented by the corresponding sid
func (pder *Provider) SessionRead(sid string) (ISession, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*Session), nil
	}
	sess, err := pder.SessionInit(sid)
	return sess, err
}

// SessionDestroy : given an sid, deletes the corresponding session
func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

// SessionGC : deletes expired session variables according to maxLifeTime
func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*Session).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*Session).sid)
		} else {
			break
		}
	}
}

// SessionUpdate :
func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*Session).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

var pder = &Provider{list: list.New()}

// Init : initialize provider
func Init() {
	pder.sessions = make(map[string]*list.Element, 0)
	Register("memory", pder)
}
