package main

import (
	"reflect"
	"testing"
)

func Test_output(t *testing.T) {
	type args struct {
		cmds []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		want1   []byte
		wantErr bool
	}{
		{
			"0",
			args{
				cmds: []string{"ECHO", "hello"},
			},
			[]string{},
			[]byte("$5\r\nhello\r\n"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := output(tt.args.cmds)
			if (err != nil) != tt.wantErr {
				t.Errorf("output() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("output() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("output() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
