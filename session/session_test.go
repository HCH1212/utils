package session

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var sessionManager *SessionManager

func sessionHandler(w http.ResponseWriter, r *http.Request) {
	session := sessionManager.GetOrCreateSession(w, r)

	// 设置一个 session 值
	session.Set("user", "golang_dev")

	// 获取 session 值
	if user, exists := session.Get("user"); exists {
		fmt.Fprintf(w, "Hello, %v!", user)
	} else {
		fmt.Fprintln(w, "No user found in session.")
	}
}

func Test(t *testing.T) {
	// 创建一个 session 管理器，session 生命周期为 30 分钟，清理周期为 10 分钟
	sessionManager = NewSessionManager(defaultLifetime, 10*time.Minute)

	http.HandleFunc("/", sessionHandler)
	fmt.Println("Server started at http://localhost:8080")
	_ = http.ListenAndServe(":8080", nil)
}
