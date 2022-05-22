package cors

// ref: https://help.hcltechsw.com/domino/10.0.1/conf_cors_json.html

type Config struct {
	// Must be equal to "1.0".
	Version string `naming:"snake_case"`

	// Must be an array of valid CORS rules.
	Rules []*Rule `naming:"snake_case"`
}

type Rule struct {
	// Describes how to match this rule with a request URL.
	Resource *Resource `naming:"snake_case"`

	// Specifies the list of allowed origins for this resource.
	AllowOrigins []string `naming:"snake_case"`

	// Specifies the list of allowed methods for this resource and origin.
	AllowMethods []string `naming:"snake_case"`

	// When true, the CORS filter allows credentials for this resource and origin.
	AllowCredentials bool `naming:"snake_case"`

	// Specifies the list of response headers to expose to XHR clients.
	ExposeHeaders []string `naming:"snake_case"`
}

type Resource struct {
	// The resource path. Without startsWith or exact values, the CORS filter matches any request with a path that contains this value.
	Path string `naming:"snake_case"`

	// When true, the CORS filter matches only requests with a path that starts with the value of path.
	StartsWith bool `naming:"snake_case"`

	// When true, the CORS filter matches only requests with a path that is the exact value of path.
	Exact bool `naming:"snake_case"`
}
