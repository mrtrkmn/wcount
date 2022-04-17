package main

import (
	"reflect"
	"testing"
)

func Test_constructMap(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{name: "Unique values map", args: args{data: []string{"a", "b", "c"}}, want: map[string]int{"a": 1, "b": 1, "c": 1}},
		{name: "Double a value in map", args: args{data: []string{"a", "b", "c", "a"}}, want: map[string]int{"a": 2, "b": 1, "c": 1}},
		{name: "Double a and b values in map", args: args{data: []string{"a", "b", "c", "a", "b"}}, want: map[string]int{"a": 2, "b": 2, "c": 1}},
		{name: "Double a, b and c values in map", args: args{data: []string{"a", "b", "c", "a", "b", "c"}}, want: map[string]int{"a": 2, "b": 2, "c": 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructMap(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cleanData(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "Test without prefix/suffix", args: args{data: "a b c"}, want: []string{"a", "b", "c"}},
		{name: "Test with prefix", args: args{data: " a b c a"}, want: []string{"a", "b", "c", "a"}},
		{name: "Test with suffix", args: args{data: "a b c a b "}, want: []string{"a", "b", "c", "a", "b"}},
		{name: "Test with prefix and suffix", args: args{data: " a b c a b "}, want: []string{"a", "b", "c", "a", "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cleanData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_preliminaryCheck(t *testing.T) {
	type args struct {
		words int
		last  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Last = words", args: args{words: 1, last: 1}, wantErr: false},
		{name: "Last > words ", args: args{words: 1, last: 2}, wantErr: true},
		{name: "Last < words ", args: args{words: 2, last: 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := preliminaryCheck(tt.args.words, tt.args.last); (err != nil) != tt.wantErr {
				t.Errorf("preliminaryCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
