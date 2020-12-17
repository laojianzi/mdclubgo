package validator

import (
	"testing"

	"github.com/laojianzi/mdclubgo/cache"
	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/log"
)

func TestContent(t *testing.T) {
	log.Init("mdclubgo", "", true)
	defer log.Close()

	type args struct {
		to      string
		appName string
		code    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				to:      "noreply@mdclubgo.localhost",
				appName: "mdclubgo",
				code:    "000000-0",
			},
			want: `To: noreply@mdclubgo.localhost
Subject: 你正在注册 mdclubgo 账号

你正在注册 mdclubgo 账号，验证码为 000000-0
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Content(tt.args.to, tt.args.appName, tt.args.code); got != tt.want {
				t.Errorf("Content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckCode(t *testing.T) {
	if err := conf.Init(conf.TestConf); err != nil {
		t.Fatal(err)
	}

	defer log.Close()
	cache.Init()
	defer cache.Close()

	email := "laojianzi@github.com"
	code := "test-code"
	t.Run("test check code failed", func(t *testing.T) {
		exists, err := CheckCode(email, code)
		if err != nil {
			t.Fatal(err)
		}

		if exists {
			t.Fatal("exists want false; got true")
		}
	})

	t.Run("test check code successful", func(t *testing.T) {
		code = GenerateCode(email)
		defer func() {
			_ = cache.Delete(CacheKey(email))
		}()
		exists, err := CheckCode(email, code)
		if err != nil {
			t.Fatal(err)
		}

		if !exists {
			t.Fatal("exists want true; got false")
		}
	})
}
