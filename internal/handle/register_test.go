package handle_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/laojianzi/mdclubgo/api"
	"github.com/laojianzi/mdclubgo/cache"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/db"
	"github.com/laojianzi/mdclubgo/internal/database"
	"github.com/laojianzi/mdclubgo/internal/email/validator"
	"github.com/laojianzi/mdclubgo/internal/handle"
	"github.com/laojianzi/mdclubgo/internal/present"
	"github.com/laojianzi/mdclubgo/internal/register"
	"github.com/laojianzi/mdclubgo/internal/storage"
	"github.com/laojianzi/mdclubgo/log"
)

func TestRegister(t *testing.T) {
	if err := conf.Init(conf.TestConf); err != nil {
		t.Fatal(err)
	}

	defer log.Close()
	db.Init()
	defer db.Close()
	defer db.Instance().Exec(fmt.Sprintf("TRUNCATE TABLE %s", new(database.User).TableName()))
	cache.Init()
	defer cache.Close()

	conf.StorageLocal.URL = "../../public/upload"
	storage.Init()
	defer storage.Close()

	email := "laojianzi@github.com"
	code := validator.GenerateCode(email)
	defer func() {
		_ = cache.Delete(validator.CacheKey(email))
	}()

	username := "laojianzi"
	jsonBody := `{"email":"` + email + `","email_code":"` + code +
		`","username":"` + username + `","password":"test-password"}`
	req := httptest.NewRequest("POST", "/api/users", strings.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	if err := handle.Register(api.Server().NewContext(req, rec)); err != nil {
		t.Fatal(err)
	}

	if rec.Code != http.StatusOK {
		t.Fatal("response status code not match")
	}

	var result present.Data
	var data register.Present
	result.Data = &data
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}

	if data.Email != email {
		t.Fatalf("response body 'email' want: %s; got: %s", username, data.Email)
	}

	if data.Username != username {
		t.Fatalf("response body 'username' want: %s; got: %s", username, data.Username)
	}
}
