package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

// UserRateLimiter rate limiter for user
// Ref: https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g
type UserRateLimiter struct {
	users map[string]*rate.Limiter
	mu    *sync.RWMutex
	r     rate.Limit
	b     int
}

// NewUserRateLimiter return a *UserRateLimiter
// Ref: https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g
func NewUserRateLimiter(r rate.Limit, b int) *UserRateLimiter {
	u := &UserRateLimiter{
		users: make(map[string]*rate.Limiter),
		mu:    &sync.RWMutex{},
		r:     r,
		b:     b,
	}

	return u
}

// AddUser creates a new rate limiter and adds it to the users map,
// using the User address as the key
// Ref: https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g
func (i *UserRateLimiter) AddUser(userKey string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.users[userKey] = limiter

	return limiter
}

// GetLimiter returns the rate limiter for the provided user key if it exists.
// Otherwise calls AddUser to add user key to the map
// Ref: https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g
func (i *UserRateLimiter) GetLimiter(userKey string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.users[userKey]

	if !exists {
		i.mu.Unlock()
		return i.AddUser(userKey)
	}

	i.mu.Unlock()

	return limiter
}
