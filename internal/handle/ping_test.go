package handle_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/laojianzi/mdclubgo/api"
	"github.com/laojianzi/mdclubgo/internal/handle"
	"github.com/laojianzi/mdclubgo/log"
)

func TestPing(t *testing.T) {
	log.Init()
	defer log.Close()

	req := httptest.NewRequest("GET", "/ping", nil)
	rec := httptest.NewRecorder()

	if err := handle.Ping(api.Server().NewContext(req, rec)); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Fatal("response status code not match")
	}

	result := make(map[string]string)
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}

	if result["version"] != "no version" {
		t.Fatal("response body not match")
	}
}
