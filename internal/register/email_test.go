package register

import (
	"errors"
	"fmt"
	"testing"

	"github.com/laojianzi/mdclubgo/email"
)

type mockMailer struct{}

func (mockMailer) Send(to []string, msg string) error {
	return nil
}

func TestContent(t *testing.T) {
	type args struct {
		to       string
		appName  string
		username string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				to:       "noreply@mdclubgo.localhost",
				appName:  "mdclubgo",
				username: "laojianzi",
			},
			want: `To: noreply@mdclubgo.localhost
Subject: 你已成功注册 mdclubgo 账号

欢迎 laojianzi，你已成功注册 mdclubgo 账号
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Content(tt.args.to, tt.args.appName, tt.args.username); got != tt.want {
				t.Errorf("Content() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSend(t *testing.T) {
	email.SetTestMailer(new(mockMailer))
	type args struct {
		to       string
		appName  string
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "default",
			args: args{
				to:       "noreply@mdclubgo.localhost",
				appName:  "mdclubgo",
				username: "laojianzi",
			},
			wantErr: nil,
		},
		{
			name: "register email send to can't empty",
			args: args{
				to:       "",
				appName:  "mdclubgo",
				username: "laojianzi",
			},
			wantErr: fmt.Errorf("register email send to can't empty"),
		},
		{
			name: "register email send to can't empty",
			args: args{
				to:       "noreply@mdclubgo.localhost",
				appName:  "",
				username: "laojianzi",
			},
			wantErr: fmt.Errorf("register email send app name can't empty"),
		},
		{
			name: "register email send to can't empty",
			args: args{
				to:       "noreply@mdclubgo.localhost",
				appName:  "mdclubgo",
				username: "",
			},
			wantErr: fmt.Errorf("register email send username can't empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Send(tt.args.to, tt.args.appName, tt.args.username); err != tt.wantErr && !errors.Is(err, tt.wantErr) &&
				(tt.wantErr != nil && err != nil && tt.wantErr.Error() != err.Error()) {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
