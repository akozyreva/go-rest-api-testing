package rest_api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	_ "github.com/tidwall/gjson"
	"io"
	"net/http"
	b "rest-api-basics/rest_api_wrapper"
	"testing"
)

func TestGetApiReq(t *testing.T) {
	fmt.Println("Starting the application...")
	response, err := http.Get("https://httpbin.org/ip")
	b.Check200StatusCode(t, response.StatusCode)
	assert.Nil(t, err, "The HTTP request failed with error %s\n", err)
	data, _ := io.ReadAll(response.Body)
	value := gjson.Get(string(data), "origin")
	println(value.String())
	fmt.Println(string(data))
}

func TestPostApiReq(t *testing.T) {
	type myJSON struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	}
	jsonStruct := &myJSON{
		FirstName: "Nic",
		LastName:  "Ivanov",
	}
	// jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	// myJsonString := `{"firstname": "Nic", "lastname": "Raboy"}`
	jsonValue, _ := json.Marshal(jsonStruct)
	response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	assert.Nil(t, err, "The HTTP request failed with error %s\n", err)
	data, _ := io.ReadAll(response.Body)
	fmt.Println(string(data))
	result := gjson.Get(string(data), "json.firstname")
	fmt.Println(result.Type)
	assert.Equal(t, result.Value(), "Nic")

}
