package redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"

	"github.com/laojianzi/mdclubgo/log"
)

// Cache cache adapter for redis
type Cache struct {
	username  string
	password  string
	host      string
	port      string
	namespace string
	client    redis.Conn
}

// Option option redis field
type Option func(*Cache)

// WithUsername set username
func WithUsername(username string) Option {
	return func(r *Cache) {
		r.username = username
	}
}

// WithPassword set password
func WithPassword(password string) Option {
	return func(r *Cache) {
		r.password = password
	}
}

// WithHost set host
func WithHost(host string) Option {
	return func(r *Cache) {
		r.host = host
	}
}

// WithPort set port
func WithPort(port string) Option {
	return func(r *Cache) {
		r.port = port
	}
}

// WithNamespace set namespace
func WithNamespace(namespace string) Option {
	return func(r *Cache) {
		r.namespace = namespace
	}
}

// NewRedis return a opened *Cache
func NewRedis(opt ...Option) *Cache {
	r := new(Cache)
	for _, o := range opt {
		o(r)
	}

	return r.open()
}

func (r *Cache) open() *Cache {
	var err error
	r.client, err = redis.Dial("tcp", fmt.Sprintf("%s:%s", r.host, r.port),
		redis.DialUsername(r.username), redis.DialPassword(r.password))
	if err != nil {
		log.Fatal("redis open failed: %s", err.Error())
	}

	return r
}

// Get return a string value from redis get
func (r *Cache) Get(key string, defaultV string) string {
	reply, err := redis.String(r.client.Do("GET", key))
	if err != nil || reply == "" {
		return defaultV
	}

	return reply
}

// Set set a string value to redis
// should use expire if ttl > 0
func (r *Cache) Set(key, value string, ttl int) error {
	commandName := "SET"
	args := []interface{}{key}
	if ttl > 0 {
		args = append(args, fmt.Sprintf("%d", ttl))
		commandName = "SETEX"
	}

	_, err := r.client.Do(commandName, append(args, value)...)
	return err
}

// Delete remove a redis value by key
func (r *Cache) Delete(key string) error {
	_, err := r.client.Do("DEL", key)
	return err
}

// Clear delete all the keys of all the existing redis
func (r *Cache) Clear() error {
	_, err := r.client.Do("FLUSHALL")
	return err
}

// GetMultiple return a []string value from redis multiple get
func (r *Cache) GetMultiple(keys []string, defaultV string) map[string]string {
	result := make(map[string]string)

	if len(keys) == 0 {
		return result
	}

	args := make([]interface{}, len(keys))
	for i, v := range keys {
		args[i] = v
	}

	values, err := redis.Strings(redis.Values(r.client.Do("MGET", args...)))
	if err != nil {
		return result
	}

	for i := range keys {
		result[keys[i]] = values[i]
	}

	return result
}

// SetMultiple set multiple data to redis
// should use expire if ttl > 0
func (r *Cache) SetMultiple(args map[string]string, ttl int) error {
	if len(args) == 0 {
		return nil
	}

	var err error
	_, err = r.client.Do("MULTI")
	if err != nil {
		return err
	}

	defer func() {
		_, err = r.client.Do("EXEC")
	}()

	for k, v := range args {
		if setErr := r.Set(k, v, ttl); setErr != nil {
			log.Error(fmt.Errorf("cache set multiple : %w", setErr).Error())
		}
	}

	return err
}

// DeleteMultiple remove multiple redis value by keys
func (r *Cache) DeleteMultiple(keys []string) error {
	if len(keys) == 0 {
		return nil
	}

	args := make([]interface{}, len(keys))
	for i := range keys {
		args[i] = keys[i]
	}

	_, err := r.client.Do("DEL", args...)
	return err
}

// Has as from exist
func (r *Cache) Has(key string) bool {
	_, err := r.client.Do("EXISTS", key)
	return err == nil
}

// Close for cache
func (r *Cache) Close() {
	_ = r.client.Close()
}
