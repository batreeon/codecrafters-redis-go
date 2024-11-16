package main

import (
	"reflect"
	"testing"
)

func TestParserInput(t *testing.T) {
	type args struct {
		bs []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"0",
			args{
				bs: []byte(""),
			},
			[]string{},
			false,
		},
		{
			"1",
			args{
				bs: []byte("+OK\r\n"),
			},
			[]string{"OK"},
			false,
		},
		{
			"2",
			args{
				bs: []byte("$5\r\nhello\r\n"),
			},
			[]string{"hello"},
			false,
		},
		{
			"3",
			args{
				bs: []byte("$0\r\n\r\n"),
			},
			[]string{""},
			false,
		},
		{
			"4",
			args{
				bs: []byte("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"),
			},
			[]string{"hello", "world"},
			false,
		},
		{
			"4",
			args{
				bs: []byte("*2\r\n$4\r\nECHO\r\n$3\r\nhey\r\n"),
			},
			[]string{"ECHO", "hey"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParserInput(tt.args.bs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParserInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParserInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsedLength(t *testing.T) {
	type args struct {
		pattern string
		v       []any
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"0",
			args{
				pattern: simpleStrings,
				v: []any{"hello"},
			},
			8,
		},
		{
			"1",
			args{
				pattern: bulkStrings,
				v: []any{15, "helloworldhello"},
			},
			22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsedLength(tt.args.pattern, tt.args.v...); got != tt.want {
				t.Errorf("parsedLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
