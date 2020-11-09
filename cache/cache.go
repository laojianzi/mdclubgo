package cache

import (
	"fmt"
	"strings"

	"github.com/laojianzi/mdclubgo/cache/redis"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/log"
)

var instance Psr16Cache

func parseRedisHostPort(info string) (host, port string) {
	host, port = "127.0.0.1", "6379"
	if strings.Contains(info, ":") && !strings.HasSuffix(info, "]") {
		idx := strings.LastIndex(info, ":")
		host = info[:idx]
		port = info[idx+1:]
	} else if len(info) > 0 {
		host = info
	}

	return host, port
}

// Init for cache
func Init() {
	typ := strings.ToLower(strings.TrimSpace(conf.Cache.Type))

	switch typ {
	case Redis:
		host, port := parseRedisHostPort(conf.Cache.Host)
		namespace := "0"
		if conf.Cache.Namespace != "" {
			namespace = conf.Cache.Namespace
		}

		opts := []redis.Option{
			redis.WithHost(host),
			redis.WithPort(port),
			redis.WithNamespace(namespace),
		}

		if conf.Cache.Password != "" {
			opts = append(opts, redis.WithUsername(conf.Cache.Username), redis.WithPassword(conf.Cache.Password))
		}

		instance = redis.NewRedis(opts...)
	case Memcached:
		// TODO: add memcached
	case PDO:
		// TODO: add pdo
	default:
		log.Fatal(fmt.Errorf("unrecognized dialect: %s", typ).Error())
	}
}

// Psr16Cache ref php Psr16Cache
// Ref: https://github.com/symfony/symfony/blob/5.x/src/Symfony/Component/Cache/Psr16Cache.php
type Psr16Cache interface {
	Get(key, defaultV string) string
	Set(key, value string, ttl int) error
	Delete(key string) error
	Clear() error
	GetMultiple(keys []string, defaultV string) map[string]string
	SetMultiple(args map[string]string, ttl int) error
	DeleteMultiple(keys []string) error
	Has(key string) bool
	Close()
}

// Get get value in cache
func Get(key, defaultV string) string {
	return instance.Get(key, defaultV)
}

// Set set value to cache
func Set(key, value string, ttl int) error {
	return instance.Set(key, value, ttl)
}

// Delete delete value in cache
func Delete(key string) error {
	return instance.Delete(key)
}

// Clear clear all value in cache
func Clear() error {
	return instance.Clear()
}

// GetMultiple get multiple value in cache
func GetMultiple(keys []string, defaultV string) map[string]string {
	return instance.GetMultiple(keys, defaultV)
}

// SetMultiple set multiple value to cache
func SetMultiple(args map[string]string, ttl int) error {
	return instance.SetMultiple(args, ttl)
}

// DeleteMultiple delete multiple value in cache
func DeleteMultiple(keys []string) error {
	return instance.DeleteMultiple(keys)
}

// Has as exist
func Has(key string) bool {
	return instance.Has(key)
}

// Close for cache
func Close() {
	instance.Close()
	instance = nil
}
