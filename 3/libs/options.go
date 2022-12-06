package libs

//---------------------------------------------------------------------------------------
//---------------------------DO NOT MODIFY THIS FILE-------------------------------------
//---------------------------------------------------------------------------------------

type Option func(*config)

// WithBody - specifies a payload
func WithBody(body interface{}) Option {
	return func(opts *config) {
		opts.body = body
	}
}

// WithHeaders - specifies any message headers to be added
func WithHeaders(headers map[string]string) Option {
	return func(opts *config) {
		opts.headers = headers
	}
}

type config struct {
	body    interface{}
	headers map[string]string
}
