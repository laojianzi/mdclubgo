package validator

import (
	"testing"

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
