package handle_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/laojianzi/mdclubgo/api"
	"github.com/laojianzi/mdclubgo/cache"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/internal/handle"
	"github.com/laojianzi/mdclubgo/log"
	"github.com/laojianzi/mdclubgo/middleware"
)

func TestNewCaptcha(t *testing.T) {
	if err := conf.Init(); err != nil {
		t.Fatal(err)
	}

	log.Init(conf.App.Name, conf.Log.RootPath, conf.App.Debug)
	defer log.Close()
	cache.Init()
	defer cache.Close()

	req := httptest.NewRequest("POST", "/api/captchas", nil)
	rec := httptest.NewRecorder()
	ctx := api.Server().NewContext(req, rec)

	if err := handle.NewCaptcha(ctx); err != nil {
		middleware.ErrorHandler(err, ctx)
	}

	if rec.Code != http.StatusOK {
		t.Fatal(fmt.Sprintf("response status code not match \nstatus = %d", rec.Code))
	}

	body := rec.Body.Bytes()
	result := make(map[string]interface{})
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	if code, ok := result["code"].(float64); !ok || code != 0 {
		t.Fatal(fmt.Sprintf("response body [code] not match \nbody = %s", string(body)))
	}

	data, ok := result["data"].(map[string]interface{})
	if !ok {
		t.Fatal(fmt.Sprintf("response body [data] is empty \nbody = %s", string(body)))
	}

	if token, ok := data["captcha_token"].(string); !ok || token == "" {
		t.Fatal(fmt.Sprintf("response body [data.captcha_token] not match \nbody = %s", string(body)))
	}

	if image, ok := data["captcha_image"].(string); !ok || image == "" {
		t.Fatal(fmt.Sprintf("response body [data.captcha_image] not match \nbody = %s", string(body)))
	}
}
