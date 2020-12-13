package local

import (
	"fmt"
)

// Option set local field
type Option func(*Local)

// WithPathPrefix set path prefix
func WithPathPrefix(pathPrefix string) Option {
	return func(r *Local) {
		if pathPrefix == "" {
			pathPrefix = DefaultPathPrefix
		}

		if pathPrefix[len(pathPrefix)-1] != '/' {
			pathPrefix = fmt.Sprintf("%s/", pathPrefix)
		}

		r.pathPrefix = pathPrefix
	}
}
