package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	app := Server()

	req := httptest.NewRequest("GET", "/ping", nil)
	resp, _ := app.Test(req)

	if resp.StatusCode != 200 {
		t.Fatal("response status code not match")
	}

	result := make(map[string]string)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		t.Fatal(err)
	}

	if result["version"] != "no version" {
		t.Fatal("response body not match")
	}
}
