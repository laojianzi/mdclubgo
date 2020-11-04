package limiter

import (
	"sync"

	"golang.org/x/time/rate"
)

// IPRateLimiter rate limiter for ip
// Ref: https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g
type IPRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// NewIPRateLimiter return a *IPRateLimiter
// Ref: https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g
func NewIPRateLimiter(r rate.Limit, b int) *IPRateLimiter {
	i := &IPRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

// AddIP creates a new rate limiter and adds it to the ips map,
// using the IP address as the key
// Ref: https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g
func (i *IPRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)
	i.ips[ip] = limiter

	return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise calls AddIP to add IP address to the map
// Ref: https://dev.to/plutov/rate-limiting-http-requests-in-go-based-on-ip-address-542g
func (i *IPRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()

	return limiter
}
