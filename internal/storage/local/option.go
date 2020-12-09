package local

// Option set local field
type Option func(*Local)

// WithPathPrefix set path prefix
func WithPathPrefix(pathPrefix string) Option {
	return func(r *Local) {
		r.pathPrefix = pathPrefix
	}
}
