package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONConvert(t *testing.T) {
	corsCfg, err := loadCORSRulesFromFile("./cors-rules.json")
	if err != nil {
		log.Fatal(err.Error())
	}

	jsonStr, _ := json.MarshalIndent(corsCfg, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestLoadCORSRulesFromFile(t *testing.T) {
	assert := assert.New(t)

	cfg, _ := loadCORSRulesFromFile("./cors-rules.json")

	assert.Equal("1.0", cfg.Version)
	assert.Len(cfg.Rules, 1)

	rule := cfg.Rules[0]
	assert.Equal("/api/data/documents", rule.Resource.Path)

	assert.Len(rule.AllowOrigins, 2)
	assert.Contains(rule.AllowOrigins, "http://this.example.com")
	assert.Contains(rule.AllowOrigins, "http://that.example.com")

	assert.Len(rule.AllowMethods, 1)
	assert.Equal("GET", rule.AllowMethods[0])

	assert.True(rule.AllowCredentials)
}
