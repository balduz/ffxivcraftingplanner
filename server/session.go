package server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const appCookie = "ffxivcrafting-session"
const minRandom = 0
const maxRandom = ^uint64(0)

type Session struct {
	ID           uint64
	CraftingList []int
}

var cachedSessions sync.Map

func GetSession(r *http.Request) *Session {
	c, err := r.Cookie(appCookie)
	if err != nil {
		return nil
	}
	id, err := strconv.ParseUint(c.Value, 10, 64)
	if err != nil {
		log.Fatalf("GetSession(_): unknown cookie value %s, error: %s", c.Value, err)
	}
	if s, ok := cachedSessions.Load(id); ok {
		return s.(*Session)
	}
	return nil
}

func AddSessionCookie(w http.ResponseWriter) *Session {
	id := generateSessionID()
	s := &Session{
		ID:           id,
		CraftingList: []int{},
	}
	// Add ID to cookie
	expiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: appCookie, Value: fmt.Sprintf("%d", id), Expires: expiration}
	http.SetCookie(w, &cookie)
	cachedSessions.Store(id, s)
	return s
}

func generateSessionID() uint64 {
	return rand.Uint64()
}
