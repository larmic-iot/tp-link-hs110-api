package parser

import (
	"net"
	"reflect"
	"testing"
)

func TestParseIp(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    net.IP
		wantErr bool
	}{
		{name: "ip is valid", args: args{"10.0.1.20"}, want: net.IP{10, 0, 1, 20}, wantErr: false},
		{name: "ip is valid", args: args{"102.0.1.202"}, want: net.IP{102, 0, 1, 202}, wantErr: false},
		{name: "ip is valid", args: args{"0.0.0.0"}, want: net.IP{0, 0, 0, 0}, wantErr: false},
		{name: "ip is not valid", args: args{"0.0.0.x"}, want: nil, wantErr: true},
		{name: "ip is not valid", args: args{"0.0.0"}, want: nil, wantErr: true},
		{name: "ip is not valid", args: args{"xxx"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIp(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want.String()) {
				t.Errorf("ParseIp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
