package convert

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertSnakeCase(t *testing.T) {
	assert := assert.New(t)

	var input map[string]interface{}

	f, _ := os.Open("../cors-rules.json")
	json.NewDecoder(f).Decode(&input)

	svc := NewService()
	output, _ := svc.SnakeCase(context.TODO(), input)

	fmt.Println(output)

	_, ok := output["version"]
	assert.True(ok)

	_, ok = output["rules"]
	assert.True(ok)

	rules := output["rules"].([]interface{})
	assert.Len(rules, 1)

	rule := rules[0].(map[string]interface{})
	_, ok = rule["allow_credentials"]
	assert.True(ok)

	_, ok = rule["allow_methods"]
	assert.True(ok)

	// skip...
}
