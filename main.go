package main

import "fmt"

type CORSConfig struct {
	Version string
	Rules   []CORSRule
}

type CORSRule struct {
	Resource        CORSResource
	AllowOrigins    []string
	AllowMethods    []string
	AllowCredential bool
	ExposeHeaders   []string
}

type CORSResource struct {
	Path       string
	StartsWith bool
	Exact      bool
}

func main() {
	fmt.Println("Hello world")
}
