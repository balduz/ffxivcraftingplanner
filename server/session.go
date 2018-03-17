package server

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"ffxivcraftingplanner/cache"
)

const appCookie = "ffxivcrafting-session"
const minRandom = 0
const maxRandom = ^int(0)

// Session contains data about a user session.
type Session struct {
	ID           int
	CraftingList []int
}

var cachedSessions = cache.NewStorage()

func GetSession(r *http.Request) *Session {
	c, err := r.Cookie(appCookie)
	if err != nil {
		return nil
	}
	id, err := strconv.Atoi(c.Value)
	if err != nil {
		log.Fatalf("GetSession(_): unknown cookie value %s, error: %s", c.Value, err)
	}
	if s, ok := cachedSessions.Get(id); ok {
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
	cachedSessions.Set(id, s)
	return s
}

func generateSessionID() int {
	return rand.Int()
}
