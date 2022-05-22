package cors

import "encoding/json"

// ref: https://help.hcltechsw.com/domino/10.0.1/conf_cors_json.html

type Config struct {
	// Must be equal to "1.0".
	Version string

	// Must be an array of valid CORS rules.
	Rules []Rule
}

func (c *Config) MarshalJSON() ([]byte, error) {
	var tmp struct {
		Version string `json:"version"`
		Rules   []Rule `json:"rules"`
	}

	tmp.Version = c.Version
	tmp.Rules = c.Rules

	return json.Marshal(tmp)
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

func (r *Rule) MarshalJSON() ([]byte, error) {
	var tmp struct {
		Resource         Resource `json:"resource"`
		AllowOrigins     []string `json:"allow_origins"`
		AllowMethods     []string `json:"allow_methods"`
		AllowCredentials bool     `json:"allow_credentials"`
		ExposeHeaders    []string `json:"expose_headers"`
	}

	tmp.Resource = r.Resource
	tmp.AllowOrigins = r.AllowOrigins
	tmp.AllowMethods = r.AllowMethods
	tmp.AllowCredentials = r.AllowCredentials
	tmp.ExposeHeaders = r.ExposeHeaders

	return json.Marshal(tmp)
}

type Resource struct {
	// The resource path. Without startsWith or exact values, the CORS filter matches any request with a path that contains this value.
	Path string

	// When true, the CORS filter matches only requests with a path that starts with the value of path.
	StartsWith bool

	// When true, the CORS filter matches only requests with a path that is the exact value of path. The startsWith and exact properties are mutually exclusive.
	Exact bool
}

func (r *Resource) MarshalJSON() ([]byte, error) {
	var tmp struct {
		Path       string `json:"path"`
		StartsWith bool   `json:"starts_with"`
		Exact      bool   `json:"exact"`
	}

	tmp.Path = r.Path
	tmp.StartsWith = r.StartsWith
	tmp.Exact = r.Exact

	return json.Marshal(tmp)
}
