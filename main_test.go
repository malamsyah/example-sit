package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/malamsyah/sit"
)

func TestMain(t *testing.T) {
	// Ubah host ke localhost:8080 ketika run test
	os.Setenv("HOST", "http://localhost:8080")
	defer os.Setenv("HOST", "https://catfact.ninja")

	// Jalanin stub/mock server
	s := sit.NewSIT(t, "http://localhost:8080")
	s.StubFor("GET", "/fact", sit.NewStubOption().
		WillReturnJSON(map[string]interface{}{
			"data": map[string]interface{}{
				"name":   "john",
				"age":    20,
				"height": 178.5,
			},
		}, nil, http.StatusOK))

	// Jalanin server utama yg akan di test
	go RunServer()

	// Test
	resp, err := http.Get("http://localhost:9090/example")
	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	respBody := make(map[string]interface{})

	json.Unmarshal(body, &respBody)

	// Check response
	fmt.Println("Response body dari Server -> ", respBody)
}
