package cors

// ref: https://help.hcltechsw.com/domino/10.0.1/conf_cors_json.html

type Config struct {
	// Must be equal to "1.0".
	Version string

	// Must be an array of valid CORS rules.
	Rules []Rule
}

type Rule struct {
	// Describes how to match this rule with a request URL.
	Resource Resource

	// Specifies the list of allowed origins for this resource.
	AllowOrigins []string

	// Specifies the list of allowed methods for this resource and origin.
	AllowMethods []string

	// When true, the CORS filter allows credentials for this resource and origin.
	AllowCredentials bool

	// Specifies the list of response headers to expose to XHR clients.
	ExposeHeaders []string
}

type Resource struct {
	// The resource path. Without startsWith or exact values, the CORS filter matches any request with a path that contains this value.
	Path string

	// When true, the CORS filter matches only requests with a path that starts with the value of path.
	StartsWith bool

	// When true, the CORS filter matches only requests with a path that is the exact value of path. The startsWith and exact properties are mutually exclusive.
	Exact bool
}
