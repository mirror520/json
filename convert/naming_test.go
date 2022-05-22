package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/mirror520/json/cors"
)

func TestConvertNaming(t *testing.T) {
	var cfg cors.Config

	f, _ := os.Open("../cors-rules.json")
	json.NewDecoder(f).Decode(&cfg)

	newStruct := ConvertNaming(&cfg)

	jsonStr, _ := json.MarshalIndent(newStruct, "", "    ")
	fmt.Println(string(jsonStr))
}

func TestStringConvertSnakeCase(t *testing.T) {
	assert := assert.New(t)

	snake := SnakeCase("CambelCase")

	assert.Equal("cambel_case", snake)
}
