package session

// 实际使用时可使用中间件形式，如hertz-contrib里面

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"sync"
	"time"
)

const (
	sessionCookieName = "session_id"
	defaultLifetime   = 30 * time.Minute
)

type Session struct {
	ID     string
	Values map[string]interface{}
	Expiry time.Time
}

type SessionManager struct {
	sessions    map[string]*Session
	mu          sync.RWMutex
	lifetime    time.Duration //  session 的有效期
	cleanupTick time.Duration // 周期性清理过期 session 的时间间隔
}

func NewSessionManager(lifetime, cleanupTick time.Duration) *SessionManager {
	manager := &SessionManager{
		sessions:    make(map[string]*Session),
		lifetime:    lifetime,
		cleanupTick: cleanupTick,
	}

	// Periodically clean up expired sessions
	go manager.cleanupExpiredSessions()

	return manager
}

// 清理过期 session
func (sm *SessionManager) cleanupExpiredSessions() {
	ticker := time.NewTicker(sm.cleanupTick)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sm.mu.Lock()
			for id, session := range sm.sessions {
				if session.Expiry.Before(time.Now()) {
					delete(sm.sessions, id)
				}
			}
			sm.mu.Unlock()
		}
	}
}

// 创建新 session
func (sm *SessionManager) NewSession() *Session {
	// 使用 crypto/rand 生成安全的 session ID
	sessionID := sm.generateSessionID()
	session := &Session{
		ID:     sessionID,
		Values: make(map[string]interface{}),
		Expiry: time.Now().Add(sm.lifetime),
	}

	sm.mu.Lock()
	sm.sessions[sessionID] = session
	sm.mu.Unlock()

	return session
}

// 根据 sessionID 获取 session
func (sm *SessionManager) GetSession(sessionID string) (*Session, bool) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	session, exists := sm.sessions[sessionID]
	if exists && session.Expiry.After(time.Now()) {
		return session, true
	}
	return nil, false
}

// 删除 session
func (sm *SessionManager) DeleteSession(sessionID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	delete(sm.sessions, sessionID)
}

// 生成安全的 sessionID
func (sm *SessionManager) generateSessionID() string {
	// 使用 crypto/rand 生成一个随机字节数组，并将其编码为 base64 字符串
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		panic("failed to generate session ID")
	}
	return base64.URLEncoding.EncodeToString(b)
}

// 获取一个新的 Session 或者已存在的 Session
func (sm *SessionManager) GetOrCreateSession(w http.ResponseWriter, r *http.Request) *Session {
	// 从 Cookie 中获取 Session ID
	cookie, err := r.Cookie(sessionCookieName)
	var session *Session
	if err != nil || cookie.Value == "" {
		// 没有 Session ID，创建一个新的 Session
		session = sm.NewSession()
		http.SetCookie(w, &http.Cookie{
			Name:     sessionCookieName,
			Value:    session.ID,
			Path:     "/",
			HttpOnly: true,
			Expires:  session.Expiry,
		})
	} else {
		// 获取已存在的 Session
		session, _ = sm.GetSession(cookie.Value)
		if session == nil {
			// 如果 Session 过期或不存在，创建新的 Session
			session = sm.NewSession()
			http.SetCookie(w, &http.Cookie{
				Name:     sessionCookieName,
				Value:    session.ID,
				Path:     "/",
				HttpOnly: true,
				Expires:  session.Expiry,
			})
		}
	}
	return session
}

// 获取 session 中的值
func (s *Session) Get(key string) (interface{}, bool) {
	value, exists := s.Values[key]
	return value, exists
}

// 设置 session 中的值
func (s *Session) Set(key string, value interface{}) {
	s.Values[key] = value
}
