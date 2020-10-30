package api

import (
	"reflect"
	"testing"

	"github.com/laojianzi/mdclubgo/conf"
)

func TestServer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"should use defulat if name is empty",
			args{},
			defaultFiberConfig.ServerHeader,
		},
		{
			"use params 'name'",
			args{name: "name"},
			"name",
		},
	}

	for _, tt := range tests {
		fiberApp = new(App)

		t.Run(tt.name, func(t *testing.T) {
			conf.App.Name = tt.args.name

			if got := Server().server.Config().ServerHeader; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServerHeader = %v, want %v", got, tt.want)
			}
		})
	}
}
